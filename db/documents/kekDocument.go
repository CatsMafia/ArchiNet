package documents

import "time"

type KekDocument struct {
	Id          string `bson:"_id,omitempty"`
	UserId      string
	Text        string
	Lols        int64
	Date        time.Time
	Url_Image   string
	Hashtags    string
	LinksPeople string
	UserLols    []string
}
