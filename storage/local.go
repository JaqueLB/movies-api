package storage

import (
	"moviesapi/entities"
	"sync"
)

var once sync.Once

type Storage struct {
	Movies []entities.Movie
}

var storageInstance *Storage

func getInstance() *Storage {
	once.Do(func() {
		storageInstance = &Storage{}
	})

	return storageInstance
}

func List() []entities.Movie {
	storage := getInstance()
	return storage.Movies
}

func GetByID(ID int) *entities.Movie {
	storage := getInstance()
	for _, movie := range storage.Movies {
		if movie.ID == ID {
			return &movie
		}
	}
	return &entities.Movie{}
}

func UpdateByID(ID int, data *entities.Movie) bool {
	storage := getInstance()
	for index, movie := range storage.Movies {
		if movie.ID == ID {
			data.ID = ID
			storage.Movies[index] = *data
			return true
		}
	}

	return false
}

func DeleteByID(ID int) bool {
	storage := getInstance()
	var newSlice []entities.Movie
	for _, movie := range storage.Movies {
		if movie.ID != ID {
			newSlice = append(newSlice, movie)
		}
	}

	storage.Movies = newSlice
	return true
}

func Create(movie *entities.Movie) *entities.Movie {
	storage := getInstance()
	ID := len(storage.Movies) + 1
	movie.ID = ID
	storage.Movies = append(storage.Movies, *movie)
	return movie
}
