package main

import (
	"log"

	"book-postgres/database"
	"book-postgres/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allBooks", routes.AllBooks) // all books
	app.Post("/addBook", routes.AddBook)  // add book
	app.Post("/book", routes.GetBook)     // get book
	app.Put("/update", routes.Update)     // update
	app.Delete("/delete", routes.Delete)  // delete
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
