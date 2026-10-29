package categories

import (
	database "backend/dbConfig"
	models "backend/models"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func DeleteCategory(c *fiber.Ctx) error {

	id, rawCategory := c.Params("id"), c.Params("category")
	var category models.Category

	if id == "1" {
		return c.Status(404).JSON(fiber.Map{"error": "Вы не можете удалить эту категорию"})
	}

	if err := database.DB.First(&category, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Категория не найдена"})
	}

	decodedCategory, err := url.QueryUnescape(rawCategory)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка декодирования URL")
	}

	if err := database.DB.Unscoped().Delete(&models.Task{}, "category = ?", decodedCategory).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка БД"})
	}

	if err := database.DB.Unscoped().Delete(&models.Category{}, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка БД"})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
