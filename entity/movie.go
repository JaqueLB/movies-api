package entity

import "time"

type Movie struct {
	ID       int
	Name     string         `json:"name"`
	Sessions []MovieSession `json:"sessions"`
}

type MovieSession struct {
	DateTime time.Time `json:"datetime"`
	Room     int       `json:"room"`
}
