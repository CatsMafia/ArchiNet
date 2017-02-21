package documents

import "time"

type KekDocument struct {
	Id     string `bson:"_id,omitempty"`
	UserId string
	Text   string
	Rate   int64
	Date   time.Time
}
