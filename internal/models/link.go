package models

import "time"

type Link struct {
	ID           int       `json:"id"`
	Hash         string    `json:"hash"`
	Url          string    `json:"url"`
	CreationDate time.Time `json:"creation_date"`
}
