package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"

	"csrf/routes"
)

func main() {
	// Ensure that the engine is properly configured with the embedded filesystem
	engine := html.NewFileSystem(http.Dir("./views"), ".html")
	engine.Layout("embed")    // Optional. Default: "embed"
	engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	// ### NORMAL SERVER ###
	// Create a new Fiber instance with the specified views engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware setup
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Register application routes
	routes.RegisterRoutes(app)
	fmt.Printf("Server started and listening at localhost:3000 - csrfActive: %v\n", len(os.Args) > 1 && os.Args[1] == "withoutCsrf")

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
