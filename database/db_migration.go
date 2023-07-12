package database

import (
	modeluser "landtick_backend/models/user"
	"landtick_backend/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		modeluser.User{},
	)
	if err != nil {
		panic(err)
	}
}