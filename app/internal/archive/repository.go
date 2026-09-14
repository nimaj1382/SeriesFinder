package archive

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(archive *Archive) error {
	return r.db.Create(archive).Error
}

func (r *Repository) Remove(archive *Archive) error {
	return r.db.Delete(archive).Error
}

func (r *Repository) Exists(archive *Archive) (bool, error) {
	var count int64
	err := r.db.Model(&Archive{}).Where("url = ?", archive.URL).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) GetAll() ([]*Archive, error) {
	var archives []*Archive
	err := r.db.Find(&archives).Error
	if err != nil {
		return nil, err
	}
	return archives, nil
}
