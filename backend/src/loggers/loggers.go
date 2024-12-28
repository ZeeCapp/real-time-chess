package loggers

import "sync"

type ILogger interface {
	LogInfo(text string)
	LogError(text string)
}

var loggerInstance ILogger
var once sync.Once

func GetLoggerInstance() ILogger {
	once.Do(func() {
		if loggerInstance == nil {
			loggerInstance = ConsoleLogger{}
		}
	})

	return loggerInstance
}
