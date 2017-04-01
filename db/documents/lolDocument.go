package documents

import "time"

type LolDocument struct {
	Id          string `bson:"_id,omitempty"`
	UserId      string
	Text        string
	Keks        int64
	Date        time.Time
	Url_Image   string
	Hashtags    string
	LinksPeople string
	UserKeks    []string
}
