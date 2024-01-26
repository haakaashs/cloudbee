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

var File *os.File
var Generic = NewGenericLogger()

func NewGenericLogger() *GenericLogger {
	flags := log.LstdFlags | log.Lshortfile
	File, _ = os.Create("logger.log")
	return &GenericLogger{
		infoLogger:  log.New(File, "INFO", flags),
		warnLogger:  log.New(File, "WARN", flags),
		errorLogger: log.New(File, "ERROR", flags),
		fatalLogger: log.New(File, "FATAL", flags),
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
