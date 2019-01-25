package logger

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writeList []LogWriter
}

//Register  注册
func (l *Logger) Register(writer LogWriter) {
	l.writeList = append(l.writeList, writer)
}

//Log 显示
func (l *Logger) Log(data interface{}) {
	for _, write := range l.writeList {
		_ = write.Write(data)
	}
}

//NewLogger 创建
func NewLogger() *Logger {
	return &Logger{}
}
