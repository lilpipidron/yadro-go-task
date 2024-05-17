package events

import (
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/club"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
	"github.com/lilpipidron/yadro-go-task/internal/table"
)

type Twelfth struct {
	Time time.Time
	Name string
}

func (event *Twelfth) Execution(club *club.Club, log lg.Log) {
	tableID := <-club.EvalibleTables
	log.Println(event.Time, 12, event.Name, tableID)

	client := club.Clients[event.Name]
	table := table.Table{Client: client, ID: tableID}
	club.Tables[tableID] = table
}
