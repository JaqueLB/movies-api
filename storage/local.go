package storage

import (
	"moviesapi/entity"
	"moviesapi/external"
	"sync"
)

var once sync.Once

type IStorage interface {
	List() []entity.Movie
	GetByID(ID int) *entity.Movie
	UpdateByID(ID int, data *external.MovieRequest) bool
	DeleteByID(ID int) bool
	Create(data *external.MovieRequest) *entity.Movie
}

type LocalStorage struct {
	Movies []entity.Movie
}

var storageInstance *LocalStorage

func GetInstance(instance *LocalStorage) *LocalStorage {
	once.Do(func() {
		storageInstance = instance
	})

	return storageInstance
}

func (storage *LocalStorage) List() []entity.Movie {
	return storage.Movies
}

func (storage *LocalStorage) GetByID(ID int) *entity.Movie {
	for _, movie := range storage.Movies {
		if movie.ID == ID {
			return &movie
		}
	}
	return &entity.Movie{}
}

func (storage *LocalStorage) UpdateByID(ID int, data *external.MovieRequest) bool {
	for index, movie := range storage.Movies {
		if movie.ID == ID {
			data.ID = ID
			storage.Movies[index] = *external.NewMovie(data)
			return true
		}
	}

	return false
}

func (storage *LocalStorage) DeleteByID(ID int) bool {
	var newSlice []entity.Movie
	for _, movie := range storage.Movies {
		if movie.ID != ID {
			newSlice = append(newSlice, movie)
		}
	}

	storage.Movies = newSlice
	return true
}

func (storage *LocalStorage) Create(data *external.MovieRequest) *entity.Movie {
	data.ID = len(storage.Movies) + 1
	movie := external.NewMovie(data)
	storage.Movies = append(storage.Movies, *movie)
	return movie
}
