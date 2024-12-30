package main

// thanks to https://github.com/Learn-by-doing/csrf-examples
import (
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"net/http"
	"os"

	"csrf/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsFS embed.FS

func main() {
	engine := html.NewFileSystem(http.FS(viewsFS), ".html")

	go func() {
		// ### EVIL SERVER ###
		// Fiber instance
		app := fiber.New(fiber.Config{Views: engine})
		app.Get("/", func(c *fiber.Ctx) error {
			// Render index - start with views directory
			return c.Render("views/layouts/main", fiber.Map{
				"Title": "Hello, World!",
			})
		})
		routes.RegisterEvilRoutes(app)
		fmt.Println("\"Evil\" server started and listening at localhost:3001")
		// Start server
		log.Fatal(app.Listen(":3001"))
	}()

	// ### NORMAL SERVER ###
	// Fiber instance
	app := fiber.New(fiber.Config{Views: engine})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index - start with views directory
		return c.Render("views/layouts/main", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	routes.RegisterRoutes(app)
	fmt.Printf("Server started and listening at localhost:3000 - csrfActive: %v\n", len(os.Args) > 1 && os.Args[1] == "withoutCsrf")
	// Start server
	log.Fatal(app.Listen(":3000"))
}
