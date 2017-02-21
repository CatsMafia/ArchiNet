package models

import (
	"github.com/CatsMafia/ArchiNet/utils"
	"time"
)

type Post struct {
	Text   string
	Rate   int64
	Id     string
	UserId string
	Date   time.Time
}

func NewPost(txt string, userId string) *Post {
	return &Post{txt, 0, utils.GenerateId(false), userId, time.Now()}
}
