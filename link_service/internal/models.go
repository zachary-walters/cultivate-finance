package services

import (
	"time"
)

type Input struct {
	Link     string `json:"link"`
	Metadata map[string]interface{}
}

type Link struct {
	Slug  string     `json:"slug" db:"slug"`
	Date  *time.Time `json:"date" db:"date"`
	InUse bool       `json:"in_use" db:"in_use"`
	Input *string    `json:"input" db:"input"`
	Link  *string    `json:"link" db:"link"`
}
