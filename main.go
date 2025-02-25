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
	engine.AddFunc("GetBaseData", func(key string) any {
		data := map[string]any{
			"title":    "Nyamuchi_",
			"username": "nyamuchi",
			"email":    "i@nyamuchi.com",
			"img":      "https://nyamuchi.com/_astro/avatar.hlvXZsLg_Z22HApl.webp",
			"motto":    "莫愁前路无知己",
			"navtab": []fiber.Map{
				{"name": "首页", "link": "/"},
				{"name": "标签", "link": "/tags"},
				{"name": "关于", "link": "/about"},
				{"name": "友链", "link": "/friends"},
				{"name": "联系我", "link": "/contactme"},
			},
			"aboutsite": "分享Web开发经验与技术见解，记录学习成长之路。",
			"social": []fiber.Map{
				{"name": "BILIBILI", "link": "https://bilibili.com", "icon": "fa-brands fa-bilibili"},
				{"name": "GITHUB", "link": "https://github.com", "icon": "fa-brands fa-github"},
				{"name": "Twitter/X", "link": "https://x.com", "icon": "fa-brands fa-x"},
				{"name": "Discord", "link": "https://discord.com", "icon": "fa-brands fa-discord"},
			},
		}
		return data[key]
	})
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	router.Setup(app)
	middleware.StaticSetup(app)
	log.Fatal(app.Listen(":3000"))
}
