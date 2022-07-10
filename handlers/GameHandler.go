package handlers

import (
	"UniWiseAssignment/data"

	"github.com/labstack/echo/v4"

	"net/http"

	//"encoding/json"
	"strconv"
)

type GamesHandler struct {
	r data.IGamesRepository
}

func NewGamesHandler(repository data.IGamesRepository) *GamesHandler {
	return &GamesHandler{r: repository}
}

func (h *GamesHandler) HandleGetGames(c echo.Context) error {
	res, err := h.r.GetAllGames()
	if err != nil {
		panic(err)
	}
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func (h *GamesHandler) HandleGetGameByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	res, err := h.r.GetGameByID(id)
	if err != nil {
		panic(err)
	}
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func (h *GamesHandler) HandleGetAllGamesByPublisher(c echo.Context) error {
	publisher := c.Param("publisher")
	res, err := h.r.GetAllGamesByPublisher(publisher)
	if err != nil {
		panic(err)
	}
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func (h *GamesHandler) HandleAddNewGame(c echo.Context) error {
	game := new(data.Game)
	if err := c.Bind(game); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, game)
}

func (h *GamesHandler) HandleEditGame(c echo.Context) error {
	return nil
}

func (h *GamesHandler) HandleRemoveGame(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	err = h.r.RemoveGame(id)
	if err != nil {
		panic(err)
	}
	return c.NoContent(http.StatusOK)
}
