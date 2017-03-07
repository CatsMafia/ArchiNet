package documents

type SessionDocument struct {
	Id   string `bson:"_id,omitempty"`
	Name string
}
