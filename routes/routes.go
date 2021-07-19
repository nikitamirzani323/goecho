package routes

import (
	"goecho/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	app.GET("/pegawai", controller.FetchAll_mpegawai)
	app.POST("/pegawai", controller.Store_mpegawai)
	app.PUT("/pegawai", controller.Update_mpegawai)
	app.DELETE("/pegawai", controller.Update_mpegawai)
	app.GET("/generate-hash/:password", controller.GenerateHashPassword)
	return app
}
