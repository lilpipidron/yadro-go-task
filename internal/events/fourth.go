package events

import (
	"strings"
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/club"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

type Fourth struct {
	Time time.Time
	Name string
}

func (event *Fourth) Execution(log lg.Log, club *club.Club) {
	log.Println(event.Time, 4, event.Name)

	club.CurrentTime = event.Time
	client, ok := club.Clients[event.Name]
	if !ok {
		th := &Thirteenth{Time: event.Time, Error: ClientUnknown}
		th.Execution(log)
		return
	}

	delete(club.Clients, event.Name)

	delete(club.Tables, client.Table)

	club.EvalibleTables[client.Table] = client.Table
}

func (event *Fourth) Parse(str string, log lg.Log) Event {
	params := strings.Split(str, " ")

	if len(params) != 3 {
		log.Fatal("incorrent number params ", str)
	}

	var err error

	if event.Time, err = time.Parse("15:04", params[0]); err != nil {
		log.Fatal("can't parser time: ", err, str)
	}

	event.Name = params[2]

	return event
}
