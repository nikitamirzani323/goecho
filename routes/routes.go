package routes

import (
	"goecho/controller"
	"goecho/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	app.GET("/pegawai", controller.FetchAll_mpegawai, middleware.IsAuthenticated)
	app.POST("/pegawai", controller.Store_mpegawai, middleware.IsAuthenticated)
	app.PUT("/pegawai", controller.Update_mpegawai, middleware.IsAuthenticated)
	app.DELETE("/pegawai", controller.Update_mpegawai, middleware.IsAuthenticated)
	app.GET("/generate-hash/:password", controller.GenerateHashPassword)

	app.GET("/generate-hash", controller.GenerateHashPassword)
	app.POST("/login", controller.CheckLogin)

	app.GET("/test-struct-validation", controller.TestStructValidation)
	app.GET("/test-variabel-validation", controller.TestVariabelValidation)
	return app
}
