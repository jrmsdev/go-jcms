package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	VERBOSE
	WARNING
	ERROR
	QUIET
	PANIC // this must be the last (higher) one
)

const defaultLevel = WARNING
const timefmt = "02/Jan/2006:15:04:05 -0700"

var outfh *os.File
var curlevel = defaultLevel

type Logger struct {
	tag         string
	tstamp      bool
	exitOnPanic bool
}

func New(tag string) *Logger {
	var err error
	if outfh == nil {
		outfh, err = os.OpenFile(
			os.DevNull,
			os.O_WRONLY,
			os.ModePerm,
		)
		if err != nil {
			panic(fmt.Sprintf("logger open: %s", err))
		}
	}
	return &Logger{tag, true, true}
}

func File(fh *os.File) error {
	if outfh != nil {
		err := outfh.Close()
		outfh = nil
		if err != nil {
			return err
		}
	}
	outfh = fh
	return nil
}

func Close() error {
	if outfh != nil {
		err := outfh.Close()
		outfh = nil
		return err
	}
	return nil
}

func Level(lvl LogLevel) {
	curlevel = lvl
}

type logEntry struct {
	lvl    LogLevel
	tag    string
	lvltag string
	tstamp bool
}

func (e *logEntry) String() string {
	s := ""
	if e.tstamp {
		s = fmt.Sprintf("[%s] ",
			time.Now().Local().Format(timefmt))
	}
	if e.lvltag != "" {
		s = fmt.Sprintf("%s[%s] ", s, e.lvltag)
	}
	return fmt.Sprintf("%s%s", s, e.tag)
}

func (l *Logger) logEntry(lvl LogLevel, lvltag string) *logEntry {
	return &logEntry{
		lvl,
		l.tag,
		lvltag,
		l.tstamp,
	}
}

func (l *Logger) log(e *logEntry, fmtstr string, args ...interface{}) {
	if e.lvl < curlevel {
		return
	}
	s := fmt.Sprintf("%s: %s\n", e.String(), fmt.Sprintf(fmtstr, args...))
	_, err := outfh.WriteString(s)
	if err != nil {
		panic(fmt.Sprintf("logger write: %s", err))
	}
}

func (l *Logger) D(fmtstr string, args ...interface{}) {
	l.log(l.logEntry(DEBUG, "D"), fmtstr, args...)
}

func (l *Logger) E(fmtstr string, args ...interface{}) {
	l.log(l.logEntry(ERROR, "E"), fmtstr, args...)
}

func (l *Logger) V(fmtstr string, args ...interface{}) {
	l.log(l.logEntry(VERBOSE, ""), fmtstr, args...)
}

func (l *Logger) W(fmtstr string, args ...interface{}) {
	l.log(l.logEntry(WARNING, "W"), fmtstr, args...)
}

func (l *Logger) Panic(fmtstr string, args ...interface{}) {
	l.log(l.logEntry(PANIC, "PANIC"), fmtstr, args...)
	if l.exitOnPanic {
		os.Exit(9)
	}
}
