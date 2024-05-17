package events

import (
	"github.com/lilpipidron/yadro-go-task/internal/club"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

type Event interface {
	Execution(log lg.Log, club *club.Club)
	Parse(str string, log lg.Log) Event
}
