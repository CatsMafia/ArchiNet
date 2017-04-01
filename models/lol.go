package models

import (
	"time"
)

type Lol struct {
	Id          string
	UserId      string
	Text        string
	Lols        int64
	Date        time.Time
	Url_Image   string
	Hashtags    string
	LinksPeople string
	UserKeks    []string
}
