package logger

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/MorselShogiew/ResizePhoto/logger/impl/std"
	"github.com/MorselShogiew/ResizePhoto/logger/opt"
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	// Fatal writes log message with fatal level and os.Exit(1) after
	Fatal(v ...interface{})
	Close()
}

type logger struct {
	kl Logger
	sl Logger
	tl Logger
}

func New(opts *opt.LoggerOpts) Logger {
	l := &logger{}

	if !opts.StdLoggerOpts.Disabled {
		l.sl = std.NewLogger(opts.StdLoggerOpts, opts.Opts)
	}

	return l
}

func (l *logger) Close() {
	if l.sl != nil {
		l.sl.Close()
	}
	if l.kl != nil {
		l.kl.Close()
	}
	if l.tl != nil {
		l.tl.Close()
	}
}

func (l *logger) Debug(v ...interface{}) {
	msg := concat(v...)
	if l.sl != nil {
		l.sl.Debug(msg)
	}
	if l.kl != nil {
		l.kl.Debug(msg)
	}
	if l.tl != nil {
		l.tl.Debug(msg)
	}
}

func (l *logger) Info(v ...interface{}) {
	msg := concat(v...)
	if l.sl != nil {
		l.sl.Info(msg)
	}
	if l.kl != nil {
		l.kl.Info(msg)
	}
	if l.tl != nil {
		l.tl.Info(msg)
	}
}

func (l *logger) Warning(v ...interface{}) {
	msg := concat(v...)
	if l.sl != nil {
		l.sl.Warning(msg)
	}
	if l.kl != nil {
		l.kl.Warning(msg)
	}
	if l.tl != nil {
		l.tl.Warning(msg)
	}
}

func (l *logger) Error(v ...interface{}) {
	msg := concat(v...)
	if l.sl != nil {
		l.sl.Error(msg)
	}
	if l.kl != nil {
		l.kl.Error(msg)
	}
	if l.tl != nil {
		l.tl.Error(msg)
	}
}

func (l *logger) Fatal(v ...interface{}) {
	msg := concat(v...)
	if l.sl != nil {
		l.sl.Fatal(msg)
	}
	if l.kl != nil {
		l.kl.Fatal(msg)
	}
	if l.tl != nil {
		l.tl.Fatal(msg)
	}
	time.Sleep(2 * time.Second)
	os.Exit(1)
}

func concat(v ...interface{}) string {
	var buffer bytes.Buffer
	for i, s := range v {
		if i == len(v)-1 {
			buffer.WriteString(fmt.Sprintf("%v", s))
		} else {
			buffer.WriteString(fmt.Sprintf("%v ", s))
		}
	}
	return buffer.String()
}
