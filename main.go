package main

import (
	"bspliter/infra"
	"bspliter/presentation"
	"bspliter/usecase"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	// db connection
	db, err := sql.Open("mysql", "bspliter:bspliter@tcp(sfoijasofie:3306)/bspliter?parseTime=true")
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
	e.POST("/events", eventHandler.Create)
	e.Logger.Fatal(e.Start(":1323"))
}
