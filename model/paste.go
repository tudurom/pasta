package model

import "time"

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Paste struct {
	BaseModel
	Filename string
	UUID     string
	Data     []byte
}
