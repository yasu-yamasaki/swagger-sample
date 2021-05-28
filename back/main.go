package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/v1/pets", show)
	e.PUT("/v1/pets", put)

	e.Logger.Fatal(e.Start(":8081"))
}

type Pet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var pets = map[string]Pet{
	"0": {
		Id:   "0",
		Name: "John",
	},
}

func show(c echo.Context) error {
	var res []Pet
	for _, v := range pets {
		res = append(res, v)
	}
	return c.JSON(http.StatusOK, res)
}

func put(c echo.Context) error {
	pet := new(Pet)
	if err := c.Bind(pet); err != nil {
		return err
	}
	pets[pet.Id] = *pet
	return c.NoContent(http.StatusOK)
}
