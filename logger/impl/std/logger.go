package std

import (
	"log"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/MorselShogiew/ResizePhoto/logger/impl"
	"github.com/MorselShogiew/ResizePhoto/logger/opt"
)

const (
	dtMask     = "2006-01-02 15:04:05.0000"
	basePrefix = ""
	baseFlag   = 0
)

const (
	defaultLevel   = infoLevel
	defaultLogFile = "/var/log/service.log"
)

const (
	debugLevel   = "DEBUG"
	infoLevel    = "INFO"
	warningLevel = "WARNING"
	errLevel     = "ERROR"
	fatalLevel   = "FATAL"
)

const (
	DEBUG   = 40
	INFO    = 30
	WARNING = 20
	ERROR   = 10
	FATAL   = 0
)

type Logger struct {
	env    string
	level  int
	logger *log.Logger
}

func NewLogger(selfOpts *opt.StdLoggerOpts, opts *opt.GeneralOpts) *Logger {
	level := opts.Level
	if level == "" {
		level = defaultLevel
	}
	logFile := selfOpts.LogFile
	if logFile == "" {
		logFile = defaultLogFile
	}

	var logLevelMap = map[string]int{
		debugLevel:   DEBUG,
		infoLevel:    INFO,
		warningLevel: WARNING,
		errLevel:     ERROR,
		fatalLevel:   FATAL,
	}

	var logger *log.Logger
	if !selfOpts.Stdout {
		logger = log.New(&lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    5,
			MaxBackups: 20,
			MaxAge:     60,
		}, basePrefix, baseFlag)
	} else {
		logger = log.New(os.Stdout, basePrefix, baseFlag)
	}

	l := &Logger{
		env:    opts.Env,
		level:  logLevelMap[level],
		logger: logger,
	}

	return l
}

func (l *Logger) Close() {}

func (l *Logger) Debug(v ...interface{}) {
	now := time.Now().Format(dtMask)
	if l.level >= DEBUG {
		msg := v[0].(string)
		l.logger.Print(now+" DEBUG ", impl.GetFuncName(), " ", msg)
	}
}

func (l *Logger) Info(v ...interface{}) {
	now := time.Now().Format(dtMask)
	if l.level >= INFO {
		msg := v[0].(string)
		l.logger.Print(now+" INFO ", impl.GetFuncName(), " ", msg)
	}
}

func (l *Logger) Warning(v ...interface{}) {
	now := time.Now().Format(dtMask)
	if l.level >= WARNING {
		msg := v[0].(string)
		l.logger.Print(now+" WARNING ", impl.GetFuncName(), " ", msg)
	}
}

func (l *Logger) Error(v ...interface{}) {
	now := time.Now().Format(dtMask)
	if l.level >= ERROR {
		msg := v[0].(string)
		l.logger.Print(now+" ERROR ", impl.GetFuncName(), " ", msg)
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	now := time.Now().Format(dtMask)
	msg := v[0].(string)
	l.logger.Print(now+" FATAL ", impl.GetFuncName(), " ", msg)
}
