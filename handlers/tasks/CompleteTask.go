package tasks

import (
	database "backend/dbConfig"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func CompleteTask(c *fiber.Ctx) error {

	id := c.Params("id")

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Задача не найдена"})
	}

	if err := database.DB.Model(&task).Where("id = ?", id).Update("is_completed", !task.IsCompleted).Error; err != nil {

		return c.Status(500).JSON(fiber.Map{"error": "Не удалось обновить состояние задачи"})
	}

	return c.Status(fiber.StatusOK).JSON("Server: Состояние задачи было обновлено")

}
