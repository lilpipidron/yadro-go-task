package logger

import (
	"fmt"
	"os"
	"time"
)

type Log interface {
	Fatal(a ...interface{})
	Info(a ...interface{})
}

type logger struct{}

func NewLogger() *logger {
	return &logger{}
}

func (log *logger) Fatal(a ...interface{}) {
	t := time.Now()
	timeStr := t.Format("2006-01-02 15:04:05")

	msg := ""
	for _, arg := range a {
		msg += fmt.Sprintf("%+v ", arg)
	}

	fmt.Printf("%s \033[31m[ERROR] %s\033[0m\n", timeStr, msg)

	os.Exit(1)
}

func (log *logger) Info(a ...interface{}) {
	t := time.Now()
	timeStr := t.Format("2006-01-02 15:04:05")

	msg := ""
	for _, arg := range a {
		msg += fmt.Sprintf("%+v ", arg)
	}

	fmt.Printf("%s \033[34m[INFO] %s\033[0m\n", timeStr, msg)
}
