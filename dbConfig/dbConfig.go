package dbConfig

import (
	"backend/models"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error
	
	DB, err = gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	log.Println("Подключение к SQLite успешно установлено.")

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}

	err = DB.AutoMigrate(&models.Category{})
	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}

	uncategorized := models.Category{
		Category:    "Uncategorized",
		Description: "Uncategorized tasks group",
	}
	
	err = DB.FirstOrCreate(&uncategorized, models.Category{Category: "Uncategorized"}).Error
	if err != nil {
		log.Fatal("Ошибка при создании категории по умолчанию: ", err)
	}

	log.Println("Миграция базы данных успешно завершена.")
}

