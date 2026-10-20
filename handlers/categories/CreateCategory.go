package categories

import (
	models "backend/models"
	database "backend/dbConfig"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {

	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	if err := database.DB.Create(category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось создать категорию"})
	}

	return c.Status(fiber.StatusCreated).JSON(category)

}
