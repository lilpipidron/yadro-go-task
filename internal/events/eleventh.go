package events

import (
	"time"

	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

type Eleventh struct {
	Time time.Time
	Name string
}

func (event *Eleventh) Execution(time time.Time, name string, log lg.Log) {
	log.Println(time, 11, name)
}
