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

	app.Get("/about", func(c fiber.Ctx) error {
		return c.Render("pages/about", fiber.Map{}, "Layout/main")
	})

	app.Get("/friends", func(c fiber.Ctx) error {
		return c.Render("pages/friend", fiber.Map{}, "Layout/main")
	})

	app.Get("/post/:name", func(c fiber.Ctx) error {
		return c.Render("pages/post", fiber.Map{}, "Layout/main")
	})

	app.Get("/tags", func(c fiber.Ctx) error {
		return c.Render("tag", fiber.Map{}, "Layout/main")
	})
}
