package config

import (
	"bufio"
	"strconv"
	"strings"
	"time"

	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
)

type Config struct {
	Start time.Time
	End   time.Time
	N     int
	Cost  int
}

func MustLoad(reader *bufio.Reader, log lg.Log) *Config {
	conf := &Config{
		Start: time.Time{},
		End:   time.Time{},
		N:     0,
		Cost:  0,
	}

	amountTables, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("can't read table amount: ", err)
	}

	amountTables = strings.TrimSpace(amountTables)

	if conf.N, err = strconv.Atoi(amountTables); err != nil {
		log.Fatal("incorrect table number format: ", err)
	}

	times, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("can't read time start and time and: ", err)
	}

	times = strings.TrimSpace(times)

	startEnd := strings.Split(times, " ")

	if conf.Start, err = time.Parse("15:04", startEnd[0]); err != nil {
		log.Fatal("incorrect start time format: ", err)
	}

	if conf.End, err = time.Parse("15:04", startEnd[1]); err != nil {
		log.Fatal("incorrect end time format: ", err)
	}

	cost, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("can't read cost: ", err)
	}

	cost = strings.TrimSpace(cost)

	if conf.Cost, err = strconv.Atoi(cost); err != nil {
		log.Fatal("incorrect cost format: ")
	}

	return conf
}
