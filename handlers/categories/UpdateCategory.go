package categories

import (
	database "backend/dbConfig"
	models "backend/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateCategory(c *fiber.Ctx) error {

	id := c.Params("id")

	var category models.Category
	var oldCategory models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Категория не найдена"})
	}

	oldCategory = category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных: " + err.Error()})
	}

	if err := database.DB.Model(&models.Task{}).Where("category = ?", oldCategory.Category).Update("category", category.Category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка БД"})
	}

	if err := database.DB.Save(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось обновить категория"})
	}

	return c.Status(fiber.StatusOK).JSON("Server: Категория была обновлена")

}
