package logger

import "time"

type Logger interface {
	log(name, message string, elapsed time.Duration)
}

type RootLogger struct {
	startTime  time.Time
	fileLogger *FileLogger
}

type ChildLogger struct {
	rootLogger *RootLogger
	name       string
	loggers    []Logger
}

func New() *RootLogger {
	return &RootLogger{
		startTime:  time.Now(),
		fileLogger: newFileLogger(),
	}
}

func (l *RootLogger) Child(name string) *ChildLogger {
	l.fileLogger.registerSource(name)
	return &ChildLogger{
		rootLogger: l,
		name:       name,
		loggers: []Logger{
			newConsoleLogger(),
			l.fileLogger,
		},
	}
}

func (c *ChildLogger) Log(message string) {
	elapsed := time.Since(c.rootLogger.startTime) // Keep full precision for sorting
	for _, lg := range c.loggers {
		lg.log(c.name, message, elapsed)
	}
}

func (l *RootLogger) Flush(filename string) error {
	return l.fileLogger.flush(filename)
}
