package entities

import "time"

type Movie struct {
	ID       int
	Name     string          `json:"name"`
	Sessions []MovieSessions `json:"sessions"`
}

type MovieSessions struct {
	ExhibitionDateTime time.Time `json:"exhibition_datetime"`
	Room               int       `json:"room"`
}
