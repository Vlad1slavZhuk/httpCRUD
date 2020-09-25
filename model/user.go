package model

type User struct {
	id        uint64 `json:"id"`
	firstName string `json:"first_name"`
	lastName  string `json:"last_name"`
}
