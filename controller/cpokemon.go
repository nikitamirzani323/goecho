package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPokemon(c echo.Context) error {
	pokemonName := c.FormValue("nama")
	client := http.DefaultClient

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+pokemonName, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	bd, _ := ioutil.ReadAll(res.Body)

	return c.JSON(http.StatusOK, string(bd))
}
