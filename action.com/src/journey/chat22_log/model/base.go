package model

type LoggerUtil interface {
	Info(data interface{}) error
}

type Logger struct {
	arrayList []LoggerUtil
}

func (log *Logger) Register(lu LoggerUtil) {
	log.arrayList = append(log.arrayList, lu)
}

func (log *Logger) WriteLog(data interface{}) {
	for _, lu := range log.arrayList {
		_ = lu.Info(data)
	}
}

func NewLogger() *Logger {
	return &Logger{}
}
