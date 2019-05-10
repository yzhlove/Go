//日志监控系统
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

/*
nginx access_log日志

remote_addr : 客户端地址
remote_user : 客户端用户名
time_local : 服务器时间
request : 请求内容，包括方法名，地址，和http协议
http_host : 用户请求是使用的http地址
status : 返回的http 状态码
request_length : 请求大小
body_bytes_sent : 返回的大小
http_referer : 来源页
http_user_agent : 客户端名称
request_time : 整体请求延时
*/

type LogProcess struct {
	read      chan []byte
	write     chan *Message
	readTool  Reader
	writeTool Writer
}

//Reader 读取接口
type Reader interface {
	Read(read chan []byte)
}

//Writer 写入接口
type Writer interface {
	Write(write chan *Message)
}

type ReadFromFile struct {
	path string
}

type WriteToInfluxDB struct {
	influxDBDns string
}

//将解析到的内容存到Message
type Message struct {
	TimeLocal                    time.Time
	BytesSent                    int //流量
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

//系统监控
type SystemInfo struct {
	HandleLine   int     `json:"handle_line"`    //总处理行数
	Tps          float64 `json:"tps"`            //系统吞吐量
	ReadChanLen  int     `json:"read_chan_len"`  //读通道长度
	WriteChanLen int     `json:"write_chan_len"` //写通道长度
	RunTime      string  `json:"run_time"`       //运行总时长
	ErrNum       int     `json:"err_num"`        //错误数
}

const (
	TypeHeadLine = 0
	TypeErrum    = 1
)

var TypeMonitorChan = make(chan int, 1024)

type Monitor struct {
	StartTime time.Time
	data      SystemInfo
	tpsSli    []int
}

func (m *Monitor) start(lop *LogProcess) {
	go func() {
		for n := range TypeMonitorChan {
			switch n {
			case TypeHeadLine:
				m.data.HandleLine += 1
			case TypeErrum:
				m.data.ErrNum += 1
			}
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for {
			<-ticker.C
			m.tpsSli = append(m.tpsSli, m.data.HandleLine)
			if len(m.tpsSli) > 2 {
				m.tpsSli = m.tpsSli[1:]
			}
		}
	}()
	http.HandleFunc("/monitor", func(writer http.ResponseWriter, request *http.Request) {
		m.data.RunTime = time.Now().Sub(m.StartTime).String()
		m.data.ReadChanLen = len(lop.read)
		m.data.WriteChanLen = len(lop.write)
		if len(m.tpsSli) >= 2 {
			m.data.Tps = float64(m.tpsSli[1]-m.tpsSli[0]) / 5
		}
		ret, _ := json.MarshalIndent(m.data, "", "\t")
		_, _ = io.WriteString(writer, string(ret))
	})
	_ = http.ListenAndServe(":9193", nil)
}

//读取
func (lop *ReadFromFile) Read(read chan []byte) {
	//读取模块
	f, err := os.Open(lop.path)
	if err != nil {
		panic(fmt.Sprintf("open file erros:%s ", err.Error()))
	}

	//将文件指针移动到文件末尾
	_, _ = f.Seek(0, 2)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes Error:%s ", err.Error()))
		}
		TypeMonitorChan <- TypeHeadLine
		//去掉"\n"
		read <- line[:len(line)-1]
	}
}

/**
日志示例:
	172.0.0.12 - - [04/Mar/2018:13:49:52 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854
正则表达式:
	([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`
*/
// 解析
func (lop *LogProcess) Process() {

	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range lop.read {
		ret := r.FindStringSubmatch(string(v))
		if len(ret) != 14 {
			TypeMonitorChan <- TypeErrum
			log.Println("FindStringSubmatch fail:", string(v))
			continue
		}
		message := &Message{}
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		if err != nil {
			TypeMonitorChan <- TypeErrum
			log.Println("ParseInLocation fail:", err.Error())
			continue
		}
		message.TimeLocal = t
		byteSent, _ := strconv.Atoi(ret[8])
		message.BytesSent = byteSent

		//GET /foo?query=t HTTP/1.0
		reqSli := strings.Split(ret[6], " ")
		if len(reqSli) != 3 {
			TypeMonitorChan <- TypeErrum
			log.Println("strings.Split fail", ret[6])
			continue
		}
		message.Method = reqSli[0]
		u, err := url.Parse(reqSli[1])
		if err != nil {
			TypeMonitorChan <- TypeErrum
			log.Println("url parse fail:", err)
			continue
		}
		message.Path = u.Path
		message.Scheme = ret[5]
		message.Status = ret[7]

		upstreamTime, _ := strconv.ParseFloat(ret[12], 64)
		requestTime, _ := strconv.ParseFloat(ret[13], 64)
		message.UpstreamTime = upstreamTime
		message.RequestTime = requestTime

		lop.write <- message
	}
}

//写入模块
func (lop *WriteToInfluxDB) Write(write chan *Message) {

	infSli := strings.Split(lop.influxDBDns, "@")

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     infSli[0],
		Username: infSli[1],
		Password: infSli[2],
	})

	if err != nil {
		fmt.Println("client.NewHTTPClient fail:", err.Error())
		return
	}

	for v := range write {
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  infSli[3],
			Precision: infSli[4],
		})

		if err != nil {
			log.Fatalln(err)
		}

		tags := map[string]string{
			"Path":   v.Path,
			"Method": v.Method,
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
			log.Fatalln(err)
		}
		bp.AddPoint(pt)

		if err := c.Write(bp); err != nil {
			log.Fatalln(err)
		}
		log.Println("Write Successful")
	}

}

func main() {

	var path, influxDsn string
	flag.StringVar(&path, "path", "./access.log", "read file path")
	flag.StringVar(&influxDsn, "influxDsn", "http://127.0.0.1:8086@admin@123456@imooc@s", "influx data source")
	flag.Parse()

	r := &ReadFromFile{
		path: "./access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDns: "http://127.0.0.1:8086@admin@123456@imooc@s",
	}

	lp := &LogProcess{
		read:      make(chan []byte, 1024),
		write:     make(chan *Message, 1024),
		readTool:  r,
		writeTool: w,
	}

	go lp.readTool.Read(lp.read)
	for i := 0; i < 2; i++ {
		go lp.Process()
	}
	for i := 0; i < 4; i++ {
		go lp.writeTool.Write(lp.write)
	}
	m := &Monitor{
		StartTime: time.Now(),
		data:      SystemInfo{},
	}
	m.start(lp)

}
