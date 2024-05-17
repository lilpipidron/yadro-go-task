package main

import (
	"bufio"
	"github.com/lilpipidron/yadro-go-task/internal/club"
	"github.com/lilpipidron/yadro-go-task/internal/config"
	"github.com/lilpipidron/yadro-go-task/internal/events"
	"io"
	"os"
	"strings"

	"github.com/lilpipidron/yadro-go-task/internal/logger"
)

func main() {
	log := logger.NewLogger(os.Stdout)

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

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		str = strings.TrimSpace(str)
		params := strings.Split(str, " ")
		if params[1] == "1" {
			ev := events.First{}
			ev.Parse(str, log)
			ev.Execution(log, cl)
		} else if params[1] == "2" {
			ev := events.Second{}
			ev.Parse(str, log)
			ev.Execution(log, cl)
		} else if params[1] == "3" {
			ev := events.Third{}
			ev.Parse(str, log)
			ev.Execution(log, cl)
		} else if params[1] == "4" {
			ev := events.Fourth{}
			ev.Parse(str, log)
			ev.Execution(log, cl)
		}
	}
}
