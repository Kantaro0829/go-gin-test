package json

type JsonRequestUser struct {
	UserMail  string `json:"mail"`
	UserPassword string `json:"pass"`
	UserAge uint8   `json:"age"`
}
