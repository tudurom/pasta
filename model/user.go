package model

type User struct {
	Model
	Token   string
	IsAdmin bool
}
