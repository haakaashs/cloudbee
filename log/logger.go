package log

import (
	"log"
	"os"
)

type GenericLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

var Generic = NewGenericLogger()

func NewGenericLogger() *GenericLogger {
	flags := log.LstdFlags | log.Lshortfile
	return &GenericLogger{
		infoLogger:  log.New(os.Stdout, "INFO", flags),
		warnLogger:  log.New(os.Stdout, "WARN", flags),
		errorLogger: log.New(os.Stdout, "ERROR", flags),
		fatalLogger: log.New(os.Stdout, "FATAL", flags),
	}
}

func (g *GenericLogger) INFO(v ...interface{}) {
	g.infoLogger.Println(v...)
}

func (g *GenericLogger) WARN(v ...interface{}) {
	g.warnLogger.Println(v...)
}

func (g *GenericLogger) ERROR(v ...interface{}) {
	g.errorLogger.Println(v...)
}

func (g *GenericLogger) FATAL(v ...interface{}) {
	g.fatalLogger.Println(v...)
}
