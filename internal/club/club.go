package club

import (
	"fmt"
	"sort"
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/client"
	"github.com/lilpipidron/yadro-go-task/internal/config"
	lg "github.com/lilpipidron/yadro-go-task/internal/logger"
	"github.com/lilpipidron/yadro-go-task/internal/table"
)

type Club struct {
	Clients        map[string]client.Client
	Tables         map[int]table.Table
	EvalibleTables map[int]int
	Queue          map[string]client.Client
	Earnings       map[client.Client]time.Time
	Config         config.Config
	CurrentTime    time.Time
}

func NewClub(config config.Config) *Club {
	return &Club{
		Clients:        make(map[string]client.Client),
		Tables:         make(map[int]table.Table),
		EvalibleTables: make(map[int]int),
		Queue:          make(map[string]client.Client),
		Earnings:       make(map[client.Client]time.Time),
		Config:         config,
	}
}

func (club *Club) CalculateRevenue(log lg.Log) {
	var hours float64
	for key, val := range club.Earnings {
		fullTime := val.Sub(key.Time)
		hours += fullTime.Hours()
		minutes := hours*60 - fullTime.Minutes()

		if minutes != 0 {
			hours++
		}
	}
	fmt.Println(hours * float64(club.Config.Cost))
}

func (club *Club) CloseClub(log lg.Log) {
	var clients []string

	for val := range club.Queue {
		clients = append(clients, val)
	}

	for _, key := range club.Tables {
		client := key.Client.Name
		clients = append(clients, client)
	}

	sort.Slice(clients, func(i, j int) bool {
		return clients[i] < clients[j]
	})

	for _, name := range clients {
		if _, ok := club.Queue[name]; !ok {
			client := club.Clients[name]
			club.Earnings[client] = club.Config.End
		}
		log.Println(club.Config.End, 11, name)
	}
}
