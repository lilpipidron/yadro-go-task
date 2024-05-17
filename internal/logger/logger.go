package logger

import (
	"fmt"
	"os"
	"time"
)

type Log interface {
	Fatal(a ...interface{})
	Info(a ...interface{})
	Ptintln(a ...interface{})
}

type Logger struct {
	Out *os.File
}

func NewLogger(out *os.File) *Logger {
	return &Logger{Out: out}
}

func (log *Logger) Fatal(a ...interface{}) {
	t := time.Now()
	timeStr := t.Format(time.RFC3339)

	msg := ""
	for _, arg := range a {
		msg += fmt.Sprintf("%+v ", arg)
	}

	fmt.Fprintf(os.Stdout, "%s \033[31m[FATAL ERROR] %s\033[0m\n", []any{timeStr, msg}...)

	os.Exit(1)
}

func (log *Logger) Println(a ...interface{}) {
	msg := ""
	for _, arg := range a {
		msg += fmt.Sprintf("%+v ", arg)
	}

	fmt.Fprintln(log.Out, []any{msg}...)
}

func (log *Logger) Info(a ...interface{}) {
	t := time.Now()
	timeStr := t.Format(time.RFC3339)

	msg := ""
	for _, arg := range a {
		msg += fmt.Sprintf("%+v ", arg)
	}

	fmt.Fprintf(os.Stdout, "%s \033[34m[INFO] %s\033[0m\n", []any{timeStr, msg}...)
}
