package external

import (
	"moviesapi/entity"
	"time"
)

type MovieRequest struct {
	ID       int
	Name     string                `json:"name"`
	Sessions []MovieSessionRequest `json:"sessions"`
}

type MovieSessionRequest struct {
	DateTime string `json:"datetime"`
	Room     int    `json:"room"`
}

func NewMovie(data *MovieRequest) *entity.Movie {
	movie := &entity.Movie{
		ID:       data.ID,
		Name:     data.Name,
		Sessions: []entity.MovieSession{},
	}

	for _, session := range data.Sessions {
		movie.Sessions = append(movie.Sessions, NewMovieSession(session))
	}

	return movie
}

func NewMovieSession(data MovieSessionRequest) entity.MovieSession {
	sessionDate, err := time.Parse("02/01/2006 15:04", data.DateTime)
	if err != nil {
		return entity.MovieSession{}
	}
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return entity.MovieSession{}
	}
	return entity.MovieSession{
		DateTime: sessionDate.In(location),
		Room:     data.Room,
	}
}
