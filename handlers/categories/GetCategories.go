package categories

import "github.com/gofiber/fiber/v2"

import (
	database "backend/dbConfig"
	models "backend/models"
)

func GetCategories(c *fiber.Ctx) error {

	categories := []models.Category{}

	if err := database.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось получить задачи"})
	}

	return c.Status(fiber.StatusOK).JSON(categories)

}
