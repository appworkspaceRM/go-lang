package database

import "landtick_backend/pkg/mysql"

func RunMigration() {
	err := mysql.DB.AutoMigrate()
	if err != nil {
		panic(err)
	}
}