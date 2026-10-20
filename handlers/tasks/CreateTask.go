package tasks

import (
	database "backend/dbConfig"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	if err := database.DB.Create(task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось создать задачу"})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}
