package club

import (
	"time"

	"github.com/lilpipidron/yadro-go-task/internal/client"
	"github.com/lilpipidron/yadro-go-task/internal/config"
	"github.com/lilpipidron/yadro-go-task/internal/table"
)

type Club struct {
	Clients        map[string]client.Client
	Tables         map[int]table.Table
	EvalibleTables map[int]int
	Queue          map[string]client.Client
	Earnings       map[client.Client]time.Time
	Config         config.Config
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
