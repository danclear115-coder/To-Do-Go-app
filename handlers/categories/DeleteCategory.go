package categories

import (
	database "backend/dbConfig"
	models "backend/models"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func DeleteCategory(c *fiber.Ctx) error {

	id, rawCategory := c.Params("id"), c.Params("category")

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
