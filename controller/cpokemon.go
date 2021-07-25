package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}
type PokemonSpecies struct {
	Name string `json:"name"`
}

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
	var responseObject Response
	json.Unmarshal(bd, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))
	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}
	return c.JSON(http.StatusOK, responseObject.Name)
}
