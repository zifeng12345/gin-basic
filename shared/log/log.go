package log

import (
	"bytes"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var defaultLogger, defaultLoggerBO logrusImpl
var defaultLoggerOnce, defaultLoggerOnceBO sync.Once

// Data is
type Data struct {
	IPAddress string `` // the ip address of caller
	Session   string `` // id that generated in controller and passed to service to service for flow tracking purpose
	ActorID   string `` // could be userID from Apps or backoffice
	ActorType string `` // MOB (Mobile Apps) / BOF (Backoffice) / MSQ (message queuing) / SCH (scheduller) / SYS (system)
}

// ILogger is
type ILogger interface {
	Debug(data interface{}, description string, args ...interface{})
	Info(data interface{}, description string, args ...interface{})
	Warn(data interface{}, description string, args ...interface{})
	Error(data interface{}, description string, args ...interface{})
	Fatal(data interface{}, description string, args ...interface{})
	Panic(data interface{}, description string, args ...interface{})
	WithFile(appsName, filename string, maxAge int)
}

// LogrusImpl is
type logrusImpl struct {
	theLogger *logrus.Logger
	useFile   bool
}

func createLogger() logrusImpl {
	// formatter := logrus.JSONFormatter{}

	formatter := nested.Formatter{
		NoColors:        true,
		HideKeys:        true,
		TimestampFormat: "0102 150405.000",
		FieldsOrder:     []string{"func"},
	}

	logger := logrusImpl{theLogger: logrus.New()}
	logger.useFile = false
	logger.theLogger.SetFormatter(&formatter)

	return logger
}

// GetLog is
func GetLog() ILogger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = createLogger()
	})
	return &defaultLogger
}

func (l *logrusImpl) WithFile(appsName, filename string, maxAge int) {
	if l.useFile {
		return
	}

	if maxAge <= 0 {
		panic("maxAge should > 0")
	}

	path := filename + ".%Y%m%d.log"
	writer, _ := rotatelogs.New(
		path,
		// rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(maxAge*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(1*24)*time.Hour),
	)

	defaultLogger.theLogger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.DebugLevel: writer,
		},
		defaultLogger.theLogger.Formatter,
	))

	l.useFile = true
}

func (l *logrusImpl) getLogEntry(extraInfo interface{}) *logrus.Entry {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()

	var buffer bytes.Buffer

	buffer.WriteString("fn:")

	x := strings.LastIndex(funcName, "/")
	buffer.WriteString(funcName[x+1:])

	if extraInfo == nil {
		return l.theLogger.WithField("info", buffer.String())
	}

	data, ok := extraInfo.(Data)
	if !ok {
		return l.theLogger.WithField("info", buffer.String())
	}

	if data.IPAddress != "" {
		buffer.WriteString("|ip:")
		buffer.WriteString(data.IPAddress)
	}

	if data.Session != "" {
		buffer.WriteString("|ss:")
		buffer.WriteString(data.Session)
	}

	if data.ActorID != "" {
		buffer.WriteString("|id:")
		buffer.WriteString(data.ActorID)
	}

	if data.ActorType != "" {
		buffer.WriteString("|tp:")
		buffer.WriteString(data.ActorType)
	}

	return l.theLogger.WithField("info", buffer.String())
}

// Debug is
func (l *logrusImpl) Debug(data interface{}, description string, args ...interface{}) {
	l.getLogEntry(data).Debugf(description+"\n", args...)
}

// Info is
func (l *logrusImpl) Info(data interface{}, description string, args ...interface{}) {
	l.getLogEntry(data).Infof(description+"\n", args...)
}

// Warn is
func (l *logrusImpl) Warn(data interface{}, description string, args ...interface{}) {
	l.getLogEntry(data).Warnf(description+"\n", args...)
}

// Error is
func (l *logrusImpl) Error(data interface{}, description string, args ...interface{}) {
	l.getLogEntry(data).Errorf(description+"\n", args...)
}

// Fatal is
func (l *logrusImpl) Fatal(data interface{}, description string, args ...interface{}) {
	l.getLogEntry(data).Fatalf(description+"\n", args...)
}

// Panic is
func (l *logrusImpl) Panic(data interface{}, description string, args ...interface{}) {
	l.getLogEntry(data).Panicf(description+"\n", args...)
}
