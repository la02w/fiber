package router

import (
	"github.com/gofiber/fiber/v3"
)

var data = fiber.Map{
	"title": "",
}

func Setup(app *fiber.App) {

	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", data, "Layout/main")
	})
}
