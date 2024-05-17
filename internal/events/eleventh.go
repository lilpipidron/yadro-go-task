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

func (event *Eleventh) Execution(club *club.Club, log lg.Log) {
	log.Println(event.Time, 11, event.Name)

	client := club.Clients[event.Name]
	delete(club.Clients, event.Name)

	club.Earnings[client] = event.Time

	delete(club.Queue, event.Name)

	delete(club.Tables, client.Table)

	club.EvalibleTables[client.Table] = client.Table
}
