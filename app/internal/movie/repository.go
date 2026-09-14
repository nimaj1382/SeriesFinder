package movie

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Add(movie *Movie) error {
	return r.db.Create(movie).Error
}

func (r *Repository) Search(title string) ([]*Movie, error) {
	var movies []*Movie
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&movies).Error
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *Repository) Exists(movie *Movie) (bool, error) {
	var count int64
	err := r.db.Model(&Movie{}).Where("title = ? AND url = ?", movie.Title, movie.URL).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
