package model

import "time"

type Todo struct {
	Id          int
	Description string
	CreatedAt   time.Time
	// CreatedAt int64
	Completed bool
}
