package ecpay

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	OFFICIAL_WEBSITE_PRODUCT_TYPE = 4 //官網自售產包
)

var MERCHANT_ID string
var TRADE_DESC string
var HASH_KEY string
var HASH_IV string

func getSysParms() {
	env := os.Getenv("environment")

	switch env {
	case "ground1", "qa2":
		MERCHANT_ID = "2000132"
		TRADE_DESC = "包你發娛樂城"
		HASH_KEY = "ejCk326UnaZWKisg"
		HASH_IV = "q9jcZX8Ib9LM8wYk"
	case "pro":
		MERCHANT_ID = "3131718"
		TRADE_DESC = "包你發娛樂城"
		HASH_KEY = "DgS7R6GYAgqO3e30"
		HASH_IV = "RyR7sdwRs6KG9Mvn"
	}

}
func init() {
	getSysParms()
}

func SetEcpay(app *fiber.App) {

	app.Post("/ecpay/GetList", ECPAYList) //顯示ECPAY 儲值清單//顯示ECPAY APK儲值起單
}
