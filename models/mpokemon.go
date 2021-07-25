package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Pokemon struct {
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

func FetchAll_mpokemon(name string) (Response, error) {
	var res Response

	client := http.DefaultClient

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+name, nil)
	if err != nil {
		return res, err
	}
	resdata, err := client.Do(req)
	if err != nil {
		return res, err
	}
	bd, _ := ioutil.ReadAll(resdata.Body)
	var object Pokemon
	json.Unmarshal(bd, &object)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = object

	return res, nil
}
