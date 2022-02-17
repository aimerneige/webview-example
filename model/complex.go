package model

type Complex struct {
	User User     `json:"user"`
	Data []string `json:"data"`
}
