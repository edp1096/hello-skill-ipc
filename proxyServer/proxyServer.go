package main // import "proxyServer"

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var count = 0

func procData(data string) {
	fmt.Println("Received-", data)
}

func main() {
	go func() {
		var data string
		for {
			fmt.Scan(&data)
			procData("chan-ipc-" + data)
		}
	}()

	e := echo.New()
	e.HidePort = true
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		data := "chan-rest-" + fmt.Sprint(count)
		procData(data)
		count++
		return c.String(http.StatusOK, "Hello, World!")
	})
	go e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
