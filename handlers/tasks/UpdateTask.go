package tasks

import (
	database "backend/dbConfig"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateTask(c *fiber.Ctx) error {

	id := c.Params("id")

	var task models.Task
	
	if err := database.DB.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Задача не найден"})
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных: " + err.Error()})
	}

	if err := database.DB.Save(&task).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось обновить задачу"})
	}

	return c.Status(fiber.StatusOK).JSON("Залача была изменена")

}
