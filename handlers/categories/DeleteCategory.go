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

	if err := database.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Задача не найден"})
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
