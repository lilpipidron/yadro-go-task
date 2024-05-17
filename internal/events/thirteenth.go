package events

import (
	"time"

	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

type Thirteenth struct {
	Time  time.Time
	Error string
}

func (event *Thirteenth) Execution(log lg.Logger) {
	log.Println(event.Time, 13, event.Error)
}
