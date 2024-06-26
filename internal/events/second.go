package events

import (
	"strconv"
	"strings"
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/club"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
	"github.com/lilpipidron/yadro-go-task/internal/table"
)

type Second struct {
	Time  time.Time
	Name  string
	Table int
}

func (event *Second) Execution(log lg.Log, club *club.Club) {
	log.Println(event.Time, 2, event.Name, event.Table)

	club.CurrentTime = event.Time
	client, ok := club.Clients[event.Name]
	if !ok {
		th := &Thirteenth{Time: event.Time, Error: ClientUnknown}
		th.Execution(log)
		return
	}

	if _, ok := club.Tables[event.Table]; ok {
		th := &Thirteenth{Time: event.Time, Error: PlacesBusy}
		th.Execution(log)
		return
	}

	table := table.Table{ID: event.Table, Client: client}
	client.Table = table.ID
	client.Time = event.Time
	club.Clients[client.Name] = client
	club.Tables[event.Table] = table
	delete(club.Queue, event.Name)
	delete(club.EvalibleTables, table.ID)
}

func (event *Second) Parse(str string, log lg.Log) Event {
	params := strings.Split(str, " ")

	if len(params) != 4 {
		log.Fatal("incorrent number params ", str)
	}

	var err error

	if event.Time, err = time.Parse("15:04", params[0]); err != nil {
		log.Fatal("can't parser time: ", err, str)
	}

	event.Name = params[2]

	if event.Table, err = strconv.Atoi(params[3]); err != nil {
		log.Fatal("can't parser table id: ", err, str)
	}

	return event
}
