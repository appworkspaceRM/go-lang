package repositorydb

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func MakeRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}