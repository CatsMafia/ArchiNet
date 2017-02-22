package models

import (
	"time"
)

type Kek struct {
	Id          string
	UserId      string
	Text        string
	Rate        int64
	Date        time.Time
	Hashtags    string
	LinksPeople string
}
