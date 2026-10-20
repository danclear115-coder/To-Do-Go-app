package main

import (
	database "backend/dbConfig"
	"backend/handlers/categories"
	"backend/handlers/tasks"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	database.ConnectDB()

	api := app.Group("/api")

	api.Post("/tasks/createTask", tasks.CreateTask)
	api.Delete("/tasks/deleteTask/:id", tasks.DeleteTask)
	api.Get("/tasks/getTasks", tasks.GetTasks)
	api.Put("/tasks/updateTask/:id", tasks.UpdateTask)

	api.Post("/categories/createCategory", categories.CreateCategory)
	api.Delete("/categories/deleteCategory/:id", categories.DeleteCategory)
	api.Get("/categories/getCategories", categories.GetCategories)
	api.Put("/categories/updateCategory/:id", categories.UpdateCategory)

	log.Fatal(app.Listen(":3000"))

}
