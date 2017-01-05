package models

type Paste struct {
	Model
	Filename string
	UUID     string
	MimeType string
	Hash     uint32
}
