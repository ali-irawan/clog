package clog

import (
	"fmt"
	"io"
	stdLog "log"
	"os"
)

type StdLogger struct {
	level    int
	levelStr map[int]string
	writer   *stdLog.Logger
}

func NewStdLogger(level int, w io.Writer, prefix string, flags int) Logger {
	if w == nil {
		w = os.Stdout
	}

	l := StdLogger{
		level: level,
		levelStr: map[int]string{
			LevelPanic: "PANIC",
			LevelFatal: "FATAL",
			LevelError: "ERROR",
			LevelWarn:  "WARN",
			LevelInfo:  "INFO",
			LevelDebug: "DEBUG",
		},
		writer: stdLog.New(w, prefix, flags),
	}
	return &l
}

func (d *StdLogger) print(level int, msg string) {
	if level > d.level {
		return
	}

	d.writer.Printf("[%s] %s", d.levelStr[level], msg)
}

func (d *StdLogger) printf(level int, pattern string, args ...interface{}) {
	if level > d.level {
		return
	}

	levelStr := fmt.Sprintf("[%s] ", d.levelStr[level])

	d.writer.Printf(levelStr+pattern, args...)
}

func (d *StdLogger) printErr(level int, msg string, err error) {
	if level > d.level {
		return
	}

	levelStr := d.levelStr[level]

	d.writer.Printf("[%s] Error: %s, Desc: %s", levelStr, err, msg)
}

func (d *StdLogger) Debug(msg string) {
	d.print(LevelDebug, msg)
}

func (d *StdLogger) Debugf(format string, args ...interface{}) {
	d.printf(LevelDebug, format, args...)
}

func (d *StdLogger) Info(msg string) {
	d.print(LevelInfo, msg)
}

func (d *StdLogger) Infof(format string, args ...interface{}) {
	d.printf(LevelInfo, format, args...)
}

func (d *StdLogger) Warn(msg string) {
	d.print(LevelWarn, msg)
}

func (d *StdLogger) Warnf(format string, args ...interface{}) {
	d.printf(LevelWarn, format, args...)
}

func (d *StdLogger) Error(msg string, err error) {
	d.printErr(LevelError, msg, err)
}

func (d *StdLogger) Errorf(format string, args ...interface{}) {
	d.printf(LevelError, format, args...)
}

func (d *StdLogger) Fatal(msg string, err error) {
	d.printErr(LevelFatal, msg, err)
}

func (d *StdLogger) Fatalf(format string, args ...interface{}) {
	d.printf(LevelFatal, format, args...)
}
