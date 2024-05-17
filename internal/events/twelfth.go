package events

import (
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/club"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
	"github.com/lilpipidron/yadro-go-task/internal/table"
)

type Twelfth struct {
	Time  time.Time
	Name  string
	Table int
}

func (event *Twelfth) Execution(club *club.Club, log lg.Log) {
	log.Println(event.Time, 12, event.Name, event.Table)

	client := club.Clients[event.Name]
	client.Table = event.Table
  club.Clients[client.Name] = client
	table := table.Table{Client: client, ID: event.Table}
	club.Tables[table.ID] = table
}
