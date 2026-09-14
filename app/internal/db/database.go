package db

import (
	"TheCollector/app/internal/archive"
	"TheCollector/app/internal/movie"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("data/data.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&movie.Movie{}, &archive.Archive{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
