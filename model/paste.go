package model

type Paste struct {
	Model
	Filename  string
	FromToken string
	UUID      string
	Hash      []byte
}
