package main

import (
	"UniWiseAssignment/data"
	"UniWiseAssignment/handlers"
	"fmt"

	"database/sql"

	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	const port = 1323
	log.Print("Initializing server")
	db, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		panic(err)
	}
	rep := data.NewGamesSQLRepository(db)
	gameHandler := handlers.NewGamesHandler(rep)

	e := echo.New()

	e.GET("/games", gameHandler.HandleGetGames)
	e.GET("/games/:id", gameHandler.HandleGetGameByID)
	e.POST("/games", gameHandler.HandleAddNewGame)
	e.PUT("/games", gameHandler.HandleEditGame)
	e.DELETE("/games", gameHandler.HandleRemoveGame)
	e.GET("/publishers/:publisher", gameHandler.HandleGetAllGamesByPublisher)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	log.Printf("Server started on :%d", port)
}
