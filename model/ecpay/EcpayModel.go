package ecpay

import (
	"gorm-fiberv2-go/common/debug"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

//取的 ECPAYList
func ECPAYList(ctx *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 10000)
			c := runtime.Stack(buf, false)
			debug.ErrorLog(string(buf[:c]))
		}
	}()
	resp, err := http.Post("http://127.0.0.1:8081/ECPAYv2/product/getbymember/list", "application/json;charset=utf-8", req.Body)
	if err != nil {
		debug.ErrorLog(err.Error())
		return
	}
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(b)

}
