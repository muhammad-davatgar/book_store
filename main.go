package main

import (
	"book_store/handlers"
	"book_store/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	dsn := ""
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	b := repositories.StoreBranchRepo{DB: db}

	if err != nil {
		panic(err)
	}

	storeHandler := handlers.StoreBranchHandlers{Repo: &b}

	app := echo.New()

	app.POST("api/store_branch", storeHandler.Create)
	app.GET("/store_branch/list", storeHandler.List)
	app.Static("/", "./views")

	app.Logger.Fatal(app.Start(":8080"))
}
