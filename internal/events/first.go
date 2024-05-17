package events

import (
	"fmt"
	"strings"
	"time"

	cl "github.com/lilpipidron/yadro-go-task/internal/client"
	"github.com/lilpipidron/yadro-go-task/internal/club"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

type First struct {
	Time time.Time
	Name string
}

func (event *First) Execution(log lg.Log, club *club.Club) {
	if _, ok := club.Clients[event.Name]; ok {
		th := &Thirteenth{Time: event.Time, Error: YouShallNotPass}
		th.Execution(log)
		return
	}

	if event.Time.Before(club.Config.Start) || event.Time.After(club.Config.End) {
		// 13 event
		return
	}

	fmt.Println(event.Time, 1, event.Name)

	client := cl.Client{Name: event.Name, Table: 0}

	club.Clients[client.Name] = client
}

func (event *First) Parse(str string, log lg.Log) Event {
	params := strings.Split(str, " ")

	if len(params) != 3 {
		log.Fatal("incorrent number params")
	}

	var err error

	if event.Time, err = time.Parse("15:04", params[0]); err != nil {
		log.Fatal("can't parse time: ", err)
	}

	event.Name = params[2]

	return event
}
