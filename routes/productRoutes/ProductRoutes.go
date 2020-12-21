package productRoutes

import (
	"gorm-fiberv2-go/model/product"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {
	app.Get("/product", product.GetAllProduct)
	app.Post("/product", product.PostProduct)
	app.Post("/GetProductList", product.GetProductList) //
	app.Post("/DeleteTodo/:id", product.DeleteTodo)
	app.Post("/PatchTodo/:id", product.PatchTodo)
}
