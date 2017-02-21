package models

import (
	"github.com/CatsMafia/ArchiNet/utils"
	"time"
)

type Kek struct {
	Id     string
	UserId string
	Text   string
	Rate   int64
	Date   time.Time
}

func NewKek(txt string, userId string) *Kek {
	return &Kek{utils.GenerateId(false), userId, txt, 0, time.Now()}
}
