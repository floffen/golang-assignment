package handlers

import (
	"github.com/labstack/echo/v4"
)

type IGamesHandler interface {
	HandleGetGames(c echo.Context) error
	HandleGetGameByID(c echo.Context) error
	HandleGetAllGamesByPublisher(c echo.Context) error
	HandleAddNewGame(c echo.Context) error
	HandleEditGame(c echo.Context) error
	HandleRemoveGame(c echo.Context) error
}