package main

import (
	"os"
	"time"
)

//Monitor 监控系统

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan []byte)
}

type LogProcess struct {
	rc     chan []byte
	wc     chan []byte
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
	HandleLine   string    `json:"handleLine"`
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

}

func NewWrite(influxDsn string) (Writer, error) {

}

func NewLogProcess(reader Reader, write Writer) *LogProcess {

}

func (l *LogProcess) Process() {
}

func (m *Monitor) start(lp *LogProcess) {

}

func main() {

}
