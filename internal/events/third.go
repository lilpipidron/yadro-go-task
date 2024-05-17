package events

import (
	"strings"
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/club"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

type Third struct {
	Time time.Time
	Name string
}

func (event *Third) Execution(log lg.Log, club *club.Club) {
	log.Println(event.Time, 3, event.Name)

	if len(club.EvalibleTables) != 0 {
		th := &Thirteenth{Time: event.Time, Error: ICanWaitNoLonger}
		th.Execution(log)
	}

	if len(club.Queue) > club.Config.N {
		eleventh := &Eleventh{Time: event.Time, Name: event.Name}
		eleventh.Execution(club, log)
		return
	}
}

func (event *Third) Parse(str string, log lg.Log) Event {
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
