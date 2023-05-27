package main

import (
	"bspliter/infra"
	"bspliter/presentation"
	"bspliter/usecase"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	// db connection
	c := mysql.NewConfig()
	c.User = "bspliter"
	c.Passwd = "bspliter"
	c.Net = "tcp"
	c.Addr = "db:3306"
	c.DBName = "bspliter"
	c.Collation = "utf8mb4_0900_ai_ci"
	c.ParseTime = true
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// dependency injection
	eventRepository := infra.NewEventRepository(db)
	eventController := usecase.NewEventController(eventRepository)
	eventHandler := presentation.NewEventHandler(eventController)

	// routing
	e := echo.New()
	e.GET("/events/:id", eventHandler.Get)
	e.POST("/events", eventHandler.Create)
	e.POST("/events/:id/payments", eventHandler.AddPayment)
	e.Logger.Fatal(e.Start(":8080"))
}
