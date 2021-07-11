package main // import "proxyServer"

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func procData(source, data string) {

	if source == "rest" {
		fmt.Println(source + "-" + data)
	}

	f, err := os.OpenFile("ipc.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("ipc.log file: " + err.Error())
		return
	}

	defer f.Close()

	dt := time.Now()
	if _, err = f.WriteString(dt.Format("2006-01-02 15:04:05") + " / " + source + " / " + data + "\n"); err != nil {
		log.Println("ipc.log write: " + err.Error())
	}
}

func ipcMain() {
	var data string
	for {
		fmt.Scan(&data)
		procData("allegro", data)
	}
}

func restMain() {
	e := echo.New()
	e.HidePort = true
	e.HideBanner = true

	// Health
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "msg: OK") })

	// Command
	e.POST("/api/allegro", func(c echo.Context) error {
		var data map[string]string
		if err := c.Bind(&data); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
		}

		procData("rest", data["cmd"]+"-"+data["arg"])
		body, _ := json.Marshal(data)

		return c.String(http.StatusOK, "msg: "+string(body))
	})

	go e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}

func main() {
	go ipcMain()

	restMain()
}
