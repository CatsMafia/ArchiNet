package models

import "time"

type Shyt struct {
	Id   string
	Rate int64
	Text string
	Date time.Time
}

func NewShyt(id string, rate int64, text string, date time.Time) *Shyt {
	return &Shyt{id, rate, text, date}
}
