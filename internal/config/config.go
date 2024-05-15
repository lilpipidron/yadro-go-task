package config

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Start time.Time
	End   time.Time
	N     int
	Cost  int
}

func MustLoad(reader *bufio.Reader) *Config {
	amountTables, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("can't read table amount", err)
	}

	conf := &Config{
		Start: time.Time{},
		End:   time.Time{},
		N:     0,
		Cost:  0,
	}

	if conf.N, err = strconv.Atoi(amountTables); err != nil {
		log.Fatal("incorrect table number format", err)
	}

	times, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("can't read time start and time and", err)
	}

	startEnd := strings.Split(times, " ")

	if conf.Start, err = time.Parse("15:04", startEnd[0]); err != nil {
		log.Fatal("incorrect start time format", err)
	}

	if conf.End, err = time.Parse("15:04", startEnd[1]); err != nil {
		log.Fatal("incorrect end time format", err)
	}

	cost, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("can't read cost", err)
	}

	if conf.Cost, err = strconv.Atoi(cost); err != nil {
		log.Fatal("incorrect cost format")
	}

	return conf
}
