package model

import "time"

type Todo struct {
	Id          int
	Description string
	CreatedAt   time.Time
	Completed   bool
}
