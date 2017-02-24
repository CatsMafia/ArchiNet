package documents

type UserDocument struct {
	Id       string `bson:"_id,omitempty"`
	Username string
	Password string
}
