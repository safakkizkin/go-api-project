package migrations

import (
	"safakkizkin/config"
	models "safakkizkin/models"
)

func InitialMigration() {
	config.DB.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&models.User{})
	config.DB.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&models.Task{})
}
