package categories

import (
	models "backend/models"
	database "backend/dbConfig"
	"github.com/gofiber/fiber/v2"
)

func UpdateCategory(c *fiber.Ctx) error {

	id := c.Params("id")

	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Категория не найден"})
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных: " + err.Error()})
	}

	if err := database.DB.Save(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось обновить категория"})
	}

	return c.Status(fiber.StatusOK).JSON("Категория была изменена")

}
