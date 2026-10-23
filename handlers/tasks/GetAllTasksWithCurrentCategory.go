package tasks

import (
	database "backend/dbConfig"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllTaksWithCurrentCategory(c *fiber.Ctx) error {
	
	category := c.Params("category")
	tasks := []models.Task{}

	if err := database.DB.Find(&tasks, "category = ?", category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось получить задачи"})
	}

	return c.Status(fiber.StatusOK).JSON(tasks)
}
