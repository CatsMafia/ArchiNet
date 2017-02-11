package models

import "time"

type Shyt struct {
	Id   uint64
	Rate int64
	Text string
	Date time.Time
}

func NewShyt(id uint64, rate int64, text string, date time.Time) *Shyt {
	return &Shyt{id, rate, text, date}
}
