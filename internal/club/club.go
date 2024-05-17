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
	EvalibleTables chan int
	Queue          chan client.Client
	Earnings       map[client.Client]time.Time
	Config         config.Config
}
