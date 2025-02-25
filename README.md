# Nyamuchi

基于 Go Fiber 的轻量级博客项目

## 功能特性
- 响应式布局与暗黑模式支持
- Tailwind CSS 样式框架
- 模块化路由配置
- 静态资源中间件
- HTML 模板引擎

```tree
├── middleware
│   └── static.go
├── router
│   └── router.go
├── web
│   ├── static
│   │   └── style
│   │       └── main.css
│   └── templates
│       ├── Layout
│       │   └── main.html
│       ├── partials
│       │   ├── footer.html
│       │   └── header.html
│       ├── 404.html
│       └── index.html
├── README.md
├── go.mod
├── go.sum
└── main.go
```

```go
// Main.go
engine := html.New("web/templates", ".html")
app := fiber.New(fiber.Config{
    Views: engine,
})
router.Setup(app)
middleware.StaticSetup(app)
log.Fatal(app.Listen(":3000"))
```

```go
// Static Files
app.Use("/static", static.New("./web/static"))
// 404 Handler
app.Use(func(c fiber.Ctx) error {
    return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{}, "Layout/main")
})
```

```go
// Router
app.Get("/", func(c fiber.Ctx) error {
    return c.Render("index", fiber.Map{}, "Layout/main")
})
```

```html
<!DOCTYPE html>
<html lang="zh">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nyamuchi~</title>
    <link rel="stylesheet" href="/static/style/main.css">
</head>

<body>
    {{template "partials/header" .}}
    {{embed}}
    {{template "partials/footer" .}}
</body>

</html>
```