package events

import "time"

type Second struct {
	Time  time.Time
	Name  string
	Table int
}
