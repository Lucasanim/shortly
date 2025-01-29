package models

import "time"

type Link struct {
	ID           int       `json:"id" dynamodbav:"ID"`
	Hash         string    `json:"hash" dynamodbav:"Hash"`
	Url          string    `json:"url" dynamodbav:"Url"`
	CreationDate time.Time `json:"creation_date" dynamodbav:"CreationDate"`
}
