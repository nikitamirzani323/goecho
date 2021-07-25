package controller

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type customer struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"gte=17,lte=35"` // harus lebih dari 17 dibawah 35
}

func TestVariabelValidation(c echo.Context) error {
	v := validator.New()

	email := "bams"

	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email not valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := customer{
		Nama:   "Bams",
		Email:  "bams@gmail.com",
		Alamat: "asdasd",
		Umur:   23,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
