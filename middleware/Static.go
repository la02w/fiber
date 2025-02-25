package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func StaticSetup(app *fiber.App) {
	// Static Files
	app.Use("/static", static.New("./web/static"))
	// 404 Handler
	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{}, "Layout/main")
	})
}
