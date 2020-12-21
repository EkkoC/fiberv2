package product

import (
	"fmt"
	"gorm-fiberv2-go/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

type GetProductv2 struct {
	gorm.Model
	Serial  string `json:"Serial"`
	Name    string `json:"Name"`
	Stock   int    `json:"Stock"`
	Content string `json:"content"`
	Price   int    `json:"Price"`
	Status  int    `json:"Price"`
}

type Product struct {
	Serial  string `json:"Serial"`
	Name    string `json:"Name"`
	Stock   int    `json:"Stock"`
	Content string `json:"content"`
	Price   int    `json:"Price"`
	Status  int    `json:"Price"`
}

func MigrateProduct(sql *gorm.DB) {
	sql.AutoMigrate(&GetProductv2{})
	fmt.Println("Product Entity migrated")
}

func GetAllProduct(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	MigrateProduct(db)
	var products []GetProductv2
	db.Find(&products)
	return ctx.Status(fiber.StatusOK).JSON(products)
}

//新增Product
func PostProduct(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	MigrateProduct(db)
	type request struct {
		Serial  string `json:"Serial"`
		Name    string `json:"Name"`
		Stock   int    `json:"Stock"`
		Content string `json:"Content"`
		Price   int    `json:"Price"`
		Status  int    `json:"Price"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	particularPerson := Product{
		Serial:  body.Serial,
		Name:    body.Name,
		Stock:   body.Stock,
		Content: body.Content,
		Price:   body.Price,
		Status:  body.Status,
	}

	db.Create(&particularPerson)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":   0,
		"status": fiber.StatusCreated,
	})

}

//取的所有Product
func GetProductList(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	var products []GetProductv2
	MigrateProduct(db)
	db.Find(&products)
	return ctx.Status(fiber.StatusOK).JSON(products)
}

//取的所有Product
func GetProduct(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	MigrateProduct(db)
	var products []GetProductv2
	paramID := ctx.Params("id")
	idTodo, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}
	db.First(&products, idTodo)
	return ctx.Status(fiber.StatusOK).JSON(products)
}

//delete Product
func DeleteTodo(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	var products Product
	MigrateProduct(db)
	paramID := ctx.Params("id")
	idTodo, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}
	db.Delete(&products, idTodo)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Successfully": "Product was deleted",
	})
}

//delete Product
func PatchTodo(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	var products Product
	MigrateProduct(db)
	type request struct {
		Price   int    `json:"Price"`
		Content string `json:"Content"`
		Stock   int    `json:"Stock"`
		Name    string `json:"Name"`
		Serial  string `json:"Serial"`
		Status  int    `json:"Status"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parser JSON",
		})
	}
	paramID := ctx.Params("id")
	idTodo, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}
	db.First(&products, idTodo)
	products.Name = body.Name
	products.Serial = body.Serial
	products.Price = body.Price
	products.Content = body.Content
	products.Stock = body.Stock
	products.Status = body.Status

	db.Save(&products)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Successfully": "Todo was updated",
	})
}
