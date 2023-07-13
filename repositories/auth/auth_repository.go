package repositoryauth

import (
	modeluser "landtick_backend/models/user"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type AuthRepository interface {
	Register(user modeluser.User)(modeluser.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository)Register(user modeluser.User) (modeluser.User, error)  {
	err := r.db.Create(&user).Error
	return user, err
}