package main

import (
	"log"
	"nn/middleware"
	"nn/router"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("web/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	router.Setup(app)
	middleware.StaticSetup(app)
	log.Fatal(app.Listen(":3000"))
}
