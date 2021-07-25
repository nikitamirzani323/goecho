package controller

import (
	"goecho/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Name           string      `json:"name"`
	Order          int         `json:"order"`
	Height         int         `json:"height"`
	BaseExperience int         `json:"base_experience"`
	Weight         int         `json:"weight"`
	Abilities      []Abilities `json:"abilities"`
}
type Abilities struct {
	Slot    int     `json:"slot"`
	Ability Ability `json:"ability"`
}
type Ability struct {
	Name string `json:"name"`
	Url  string `json:url`
}

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}
type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"developer"`
}
type SensorReading struct {
	Name        string `json:"name"`
	Capacity    string `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}
type Info struct {
	Description string `json:"desc"`
}

func GetPokemon(c echo.Context) error {
	pokemonName := c.FormValue("nama")
	result, err := models.FetchAll_mpokemon(pokemonName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
	// client := http.DefaultClient

	// req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+pokemonName, nil)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	// }
	// res, err := client.Do(req)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	// }
	// bd, _ := ioutil.ReadAll(res.Body)
	// var responseObject Response
	// json.Unmarshal(bd, &responseObject)

	// author := Author{Name: "Eliiot Forber", Age: 25, Developer: true}
	// book := Book{Title: "learning python", Author: author}

	// // bytearr, err := json.Marshal(book)
	// bytearr, err := json.MarshalIndent(book, "", " ")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(bytearr))

	// jsonstring := `
	// 	{
	// 		"name": "battrey sensor",
	// 		"capacity": 40,
	// 		"time":"2019-01-21T19:07:28Z",
	// 		"info": {
	// 			"desc": "a sensor reading"
	// 		}
	// 	}`

	// var reading SensorReading
	// errr := json.Unmarshal([]byte(jsonstring), &reading)
	// if errr != nil {
	// 	fmt.Println(errr)
	// }
	// fmt.Printf("%+v\n", reading)

	// return c.JSON(http.StatusOK, responseObject)
}
