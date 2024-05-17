package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/club"
	"github.com/lilpipidron/yadro-go-task/internal/config"
	"github.com/lilpipidron/yadro-go-task/internal/events"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

func TryPlaceClient(club *club.Club, log lg.Log) {
	if len(club.EvalibleTables) > 0 && len(club.Queue) > 0 {
		var table int
		var client string
		for key := range club.Queue {
			client = key
			break
		}
		for key := range club.EvalibleTables {
			table = key
			break
		}
		tw := events.Twelfth{
			Name:  client,
			Table: table,
			Time:  club.CurrentTime,
		}
		tw.Execution(club, log)
	}
}

func main() {
	log := lg.NewLogger(os.Stdout)
	if len(os.Args) != 2 {
		log.Fatal("incorrect number of launch parameters")

	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("can't open file", err)
	}

	reader := bufio.NewReader(file)

	conf := config.MustLoad(reader, log)

	cl := club.NewClub(*conf)
	tables := make(map[int]int)
	for i := 1; i <= conf.N; i++ {
		tables[i] = i
	}
	cl.EvalibleTables = tables

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF && str == "" {
			break
		}

		str = strings.TrimSpace(str)
		params := strings.Split(str, " ")

		switch params[1] {
		case "1":
			ev := events.First{
				Time: time.Time{},
				Name: "",
			}
			ev.Parse(str, log)
			ev.Execution(log, cl)
		case "2":
			ev := events.Second{
				Time:  time.Time{},
				Name:  "",
				Table: 0,
			}
			ev.Parse(str, log)
			ev.Execution(log, cl)
		case "3":
			ev := events.Third{
				Time: time.Time{},
				Name: "",
			}
			ev.Parse(str, log)
			ev.Execution(log, cl)
		case "4":
			ev := events.Fourth{
				Time: time.Time{},
				Name: "",
			}
			ev.Parse(str, log)
			ev.Execution(log, cl)
		}
		TryPlaceClient(cl, log)
	}

	cl.CloseClub(log)
	cl.CalculateRevenue(log)
}
