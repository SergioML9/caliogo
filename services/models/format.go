package models

type Format struct {
	ID        int    `json:"id"`
	Extension string `json:"extension"`
	Mime      string `json:"mime"`
}
