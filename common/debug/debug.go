package debug

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var isdebug bool
var isTesting = "N"

func init() {
	result, err := strconv.ParseBool(os.Getenv("debug"))
	if err != nil {
		ErrorLog(err.Error())
	}
	isdebug = result
	Log(fmt.Sprintf("isdebug:%t", isdebug))
}

func Log(s string) {
	if isTesting == "Y" { //for test
		return
	}
	var text = fmt.Sprintf("[%s] [Log] %s \r\n", time.Now().Format("2006-01-02 15:04:05.00000"), s)
	var date = fmt.Sprintf("%s%s", time.Now().Format("2006-01-02"), ".txt")
	f, err := os.OpenFile(date, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Log(err.Error())
	}
	defer f.Close()
	fmt.Printf(text)
	_, err = f.WriteString(text)
	if err != nil {
		Log(err.Error())
	}
}

func ErrorLog(s string) {
	if isTesting == "Y" { //for test
		return
	}
	var text = fmt.Sprintf("[%s] [ErrorLog] %s \r\n", time.Now().Format("2006-01-02 15:04:05.00000"), s)
	var date = fmt.Sprintf("%s%s", time.Now().Format("2006-01-02"), ".txt")
	f, err := os.OpenFile(date, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Log(err.Error())
	}
	defer f.Close()
	fmt.Printf(text)
	_, err = f.WriteString(text)
	if err != nil {
		Log(err.Error())
	}
}
