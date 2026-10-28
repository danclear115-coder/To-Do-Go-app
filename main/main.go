package main

import (
	database "backend/dbConfig"
	"backend/handlers/categories"
	"backend/handlers/tasks"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	database.ConnectDB()

	api := app.Group("/api")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	api.Post("/tasks/createTask", tasks.CreateTask)
	api.Delete("/tasks/deleteTask/:id", tasks.DeleteTask)
	api.Get("/tasks/getTasks", tasks.GetTasks)
	api.Get("/tasks/getTasksByCategory/:category", tasks.GetAllTaksWithCurrentCategory)
	api.Put("/tasks/updateTask/:id", tasks.UpdateTask)

	api.Post("/categories/createCategory", categories.CreateCategory)
	api.Delete("/categories/deleteCategory/:id/:category", categories.DeleteCategory)
	api.Get("/categories/getCategories", categories.GetCategories)
	api.Put("/categories/updateCategory/:id", categories.UpdateCategory)

	log.Fatal(app.Listen(":3000"))

}
