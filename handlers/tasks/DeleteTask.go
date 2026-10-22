package tasks

import (
	database "backend/dbConfig"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteTask(c *fiber.Ctx) error {
	
	id := c.Params("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Задача не найден"})
	}

	if err := database.DB.Unscoped().Delete(&models.Task{}, "id = ?", id).Error; err != nil {
    	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка БД"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
