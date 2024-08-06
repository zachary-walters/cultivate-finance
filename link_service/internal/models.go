package services

import (
	"time"
)

type Input struct {
	Link     string `json:"link"`
	Metadata map[string]interface{}
}

type Link struct {
	Slug     string     `json:"slug" db:"slug"`
	Date     *time.Time `json:"date" db:"date"`
	Metadata *string    `json:"metadata" db:"metadata"`
	Link     *string    `json:"link" db:"link"`
}
