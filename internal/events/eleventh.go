package events

import (
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/club"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

type Eleventh struct {
	Time time.Time
	Name string
}

func (event *Eleventh) Execution(time time.Time, name string, club *club.Club, log lg.Log) {
	log.Println(time, 11, name)

	client := club.Clients[name]
	delete(club.Clients, name)

	club.Earnings[client] = time
}
