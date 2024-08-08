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

var badWords = map[string]bool{
	"anal": true,
	"anus": true,
	"asss": true,
	"boob": true,
	"cock": true,
	"cuck": true,
	"cumm": true,
	"coon": true,
	"dick": true,
	"fagg": true,
	"fagt": true,
	"fags": true,
	"fuck": true,
	"jizz": true,
	"milf": true,
	"nazi": true,
	"niga": true,
	"nigg": true,
	"nigr": true,
	"piss": true,
	"poon": true,
	"rape": true,
	"scat": true,
	"shit": true,
	"smut": true,
	"suck": true,
	"tits": true,
	"titt": true,
	"twat": true,
	"wank": true,
}
