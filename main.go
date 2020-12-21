package main

import (
	//"fiberv2/routes/productRoutes"

	"gorm-fiberv2-go/routes/productRoutes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	//設置Fiber v2以啟動服務器
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome\n")
	})
	//設置路由模型功能
	productRoutes.SetupProductRoutes(app)

	//Setting to listen server port
	app.Listen(":4000")
}
