package main

import (
	"bufio"
	"os"

	"github.com/lilpipidron/yadro-go-task/internal/config"
	"github.com/lilpipidron/yadro-go-task/internal/logger"
)

func main() {
	log := logger.NewLogger()

	if len(os.Args) != 2 {
		log.Fatal("incorrect number of launch parameters")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("can't open file", err)
	}

	reader := bufio.NewReader(file)

	conf := config.MustLoad(reader, log)

	log.Info(conf)
}
