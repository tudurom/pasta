package models

type User struct {
	Model
	Token   string
	IsAdmin bool
}
