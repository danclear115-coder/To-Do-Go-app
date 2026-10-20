package categories

import (
	models "backend/models"
	database "backend/dbConfig"
	"github.com/gofiber/fiber/v2"
)

func DeleteCategory(c *fiber.Ctx) error {

	id := c.Params("id")

	if err := database.DB.Unscoped().Delete(&models.Category{}, "id = ?", id).Error; err != nil {
    	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка БД"})
	}

	return c.SendStatus(fiber.StatusNoContent)
	
}

