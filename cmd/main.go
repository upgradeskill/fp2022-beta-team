package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"log"
	"time"
	"encoding/json"
)

func main() {
	e := echo.New()
	e.GET("/", HandlerMain)
	version := "/v1"
	e.GET(version + "/companies", HandlerCompanies)
	e.Logger.Fatal(e.Start(":8080"))
}

func HandlerMain(echCtx echo.Context) error {
	return echCtx.String(http.StatusOK, "Hello, World!")
}

type Companies struct{
	id uint64
	name string
	description string
	created_at time.Time
	updated_at time.Time
}

func HandlerCompanies(echCtx echo.Context) error {
	existingCompanies := []Companies{
		{
			id: 1,
			name: "PT Majoo Teknologi Indonesia",
			description: "Aplikasi Wirausaha",
		},
	}
	jsonCompanies, err := json.Marshal(existingCompanies)
	if err != nil {
		log.Fatal(err)
	}
	return echCtx.JSON(http.StatusOK, jsonCompanies)
}