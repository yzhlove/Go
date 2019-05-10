package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

//Monitor 监控系统

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan *Message)
}

type LogProcess struct {
	rc     chan []byte
	wc     chan *Message
	reader Reader
	writer Writer
}

type ReadFromTail struct {
	inode uint64
	fd    *os.File
	path  string
}

type WriteToInfluxDB struct {
	batch      uint16
	retry      uint8
	influxConf *InfluxConf
}

type InfluxConf struct {
	Addr, UserName, Password, Database, Precision string
}

type Message struct {
	TimeLocal                    time.Time
	BytesSent                    int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

type Monitor struct {
	listenPort string
	startTime  time.Time
	tpsSli     []int
	systemInfo SystemInfo
}

type SystemInfo struct {
	HandleLine   int       `json:"handleLine"`
	Tps          float64   `json:"tps"`
	ReadChanLen  int       `json:"readChanLen"`
	WriteChanLen int       `json:"writeChanLen"`
	RunTime      string    `json:"runTime"`
	ErrInfo      ErrorInfo `json:"errInfo"`
}

type ErrorInfo struct {
	ReadErr    int `json:"readErr"`
	ProcessErr int `json:"processErr"`
	WriteErr   int `json:"writeErr"`
}

type TypeMonitor int

const (
	TypeHandLine TypeMonitor = iota
	TypeReadErr
	TypeProcessErr
	TypeWriteErr
)

var (
	path, influxDsn, listenPort string
	processNum, writeNum        int
	TypeMonitorChan             = make(chan TypeMonitor, 200)
)

func NewReader(path string) (Reader, error) {
	var stat syscall.Stat_t
	if err := syscall.Stat(path, &stat); err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &ReadFromTail{
		inode: stat.Ino,
		fd:    f,
		path:  path,
	}, nil
}

func NewWrite(influxDsn string) (Writer, error) {
	influxDsnSli := strings.Split(influxDsn, "@")
	if len(influxDsnSli) < 5 {
		return nil, errors.New("param influxDns err")
	}
	return &WriteToInfluxDB{
		batch: 50,
		retry: 3,
		influxConf: &InfluxConf{
			Addr:      influxDsnSli[0],
			UserName:  influxDsnSli[1],
			Password:  influxDsnSli[2],
			Database:  influxDsnSli[3],
			Precision: influxDsnSli[4],
		},
	}, nil
}

func NewLogProcess(read Reader, write Writer) *LogProcess {
	return &LogProcess{
		rc:     make(chan []byte, 200),
		wc:     make(chan *Message, 200),
		reader: read,
		writer: write,
	}
}

//解析模块
func (l *LogProcess) Process() {
	rep := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range l.rc {
		TypeMonitorChan <- TypeHandLine
		ret := rep.FindStringSubmatch(string(v))
		if len(ret) < 13 {
			TypeMonitorChan <- TypeProcessErr
			log.Println("wrong input data:", v)
			continue
		}

		timeLocal, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		if err != nil {
			TypeMonitorChan <- TypeProcessErr
			log.Println("time parse error:", err)
			continue
		}

		request := ret[6]
		requestSli := strings.Split(request, " ")
		if len(requestSli) < 3 {
			TypeMonitorChan <- TypeProcessErr
			log.Println("input request wrong: ", request)
			continue
		}

		method := strings.TrimLeft(requestSli[0], "\"")
		u, err := url.Parse(requestSli[1])
		if err != nil {
			TypeMonitorChan <- TypeProcessErr
			log.Println("input url parse err: ", err)
			continue
		}

		path := u.Path
		scheme := ret[5]
		status := ret[7]
		bytesSent, _ := strconv.Atoi(ret[8])
		upstreamTime, _ := strconv.ParseFloat(ret[12], 64)
		requestTime, _ := strconv.ParseFloat(ret[13], 64)

		l.wc <- &Message{
			TimeLocal:    timeLocal,
			Path:         path,
			Method:       method,
			Scheme:       scheme,
			Status:       status,
			BytesSent:    bytesSent,
			UpstreamTime: upstreamTime,
			RequestTime:  requestTime,
		}
	}
}

//Http监听
func (m *Monitor) start(lp *LogProcess) {
	go func() {
		for n := range TypeMonitorChan {
			switch n {
			case TypeHandLine:
				m.systemInfo.HandleLine += 1
			case TypeProcessErr:
				m.systemInfo.ErrInfo.ProcessErr += 1
			case TypeReadErr:
				m.systemInfo.ErrInfo.ReadErr += 1
			case TypeWriteErr:
				m.systemInfo.ErrInfo.WriteErr += 1
			}
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for {
			<-ticker.C
			m.tpsSli = append(m.tpsSli, m.systemInfo.HandleLine)
			if len(m.tpsSli) > 2 {
				m.tpsSli = m.tpsSli[1:]
			}
		}
	}()

	http.HandleFunc("/monitor", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, m.systemStatus(lp))
	})
	log.Fatal(http.ListenAndServe(":"+m.listenPort, nil))
}

func (m *Monitor) systemStatus(lp *LogProcess) string {
	d := time.Now().Sub(m.startTime)
	m.systemInfo.RunTime = d.String()
	m.systemInfo.ReadChanLen = len(lp.rc)
	m.systemInfo.WriteChanLen = len(lp.rc)
	if len(m.tpsSli) >= 2 {
		m.systemInfo.Tps = float64(m.tpsSli[1]-m.tpsSli[0]) / 5
	}
	res, _ := json.MarshalIndent(m.systemInfo, "monitor", "\t")
	return string(res)
}

func (r *ReadFromTail) Read(rc chan []byte) {
	defer close(rc)
	var stat syscall.Stat_t
	_, _ = r.fd.Seek(0, 2)
	bf := bufio.NewReader(r.fd)

	for {
		line, err := bf.ReadBytes('\n')
		if err == io.EOF {
			if err := syscall.Stat(r.path, &stat); err != nil {
				time.Sleep(1 * time.Second)
			} else {
				nowInode := stat.Ino
				if nowInode == r.inode {
					time.Sleep(1 * time.Second)
				} else {
					r.fd.Close()
					if fd, err := os.Open(r.path); err != nil {
						panic(fmt.Sprintf("Open File err : %s", err.Error()))
					} else {
						r.fd = fd
						bf = bufio.NewReader(fd)
						r.inode = nowInode
					}
				}
			}
			continue
		} else if err != nil {
			TypeMonitorChan <- TypeReadErr
			log.Printf("readFromTail ReadBytes err: %s", err.Error())
			continue
		}
		rc <- line[:len(line)-1]
	}
}

func (w *WriteToInfluxDB) Write(wc chan *Message) {
	infClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     w.influxConf.Addr,
		Username: w.influxConf.UserName,
		Password: w.influxConf.Password,
	})
	if err != nil {
		panic(fmt.Sprintf("influxdb NewHttpClient err:%s ", err.Error()))
	}

	for {
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  w.influxConf.Database,
			Precision: w.influxConf.Precision,
		})
		if err != nil {
			panic(fmt.Sprintf("influxdb NewBatchPoints err:%s", err.Error()))
		}
		var count uint16
		for v := range wc {
			tags := map[string]string{
				"Path":   v.Path,
				"Scheme": v.Scheme,
				"Status": v.Status,
			}
			fields := map[string]interface{}{
				"UpstreamTime": v.UpstreamTime,
				"RequestTime":  v.RequestTime,
				"BytesSent":    v.BytesSent,
			}

			pt, err := client.NewPoint("nginx_log", tags, fields, v.TimeLocal)
			if err != nil {
				TypeMonitorChan <- TypeWriteErr
				log.Println("influxdb NewPoint error:", err)
				continue
			}
			bp.AddPoint(pt)
			count++
			if count > w.batch {
				break
			}
		}

		var i uint8
		for i = 1; i <= w.retry; i++ {
			if err := infClient.Write(bp); err != nil {
				TypeMonitorChan <- TypeWriteErr
				log.Printf("influxdb write err:%s retry:%d ", err.Error(), i)
				time.Sleep(1 * time.Second)
			} else {
				log.Println(w.batch, " point has written")
				break
			}
		}
	}
}

func init() {
	flag.StringVar(&path, "path", "./access.log", "log file path")
	flag.StringVar(&influxDsn, "influxDsn", "", "influxDB dsn")
	flag.StringVar(&listenPort, "listenPort", "9193", "monitor port")
	flag.IntVar(&processNum, "processNum", 1, "process goroutine num")
	flag.IntVar(&writeNum, "writeNum", 1, "write goroutine num")
	flag.Parse()
}

func main() {
	reader, err := NewReader(path)
	if err != nil {
		panic(err)
	}

	writer, err := NewWrite(influxDsn)
	if err != nil {
		panic(err)
	}

	lp := NewLogProcess(reader, writer)

	go lp.reader.Read(lp.rc)

	for i := 0; i < processNum; i++ {
		go lp.Process()
	}

	for i := 0; i < writeNum; i++ {
		go lp.writer.Write(lp.wc)
	}

	m := &Monitor{
		listenPort: listenPort,
		startTime:  time.Now(),
	}

	go m.start(lp)

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			log.Println("capture exit signal: ", s)
			os.Exit(1)
		case syscall.SIGUSR1:
			log.Println(m.systemStatus(lp))
		default:
			log.Println("capture other signal:", s)
		}
	}
}
