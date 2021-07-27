package routes

import (
	"goecho/controller"
	"goecho/middleware"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() *echo.Echo {
	app := echo.New()
	app.Use(middleware.Logger)
	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}
	app.Renderer = t
	app.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"message": "Golang Web Apps",
		})
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

	app.POST("/pokemon", controller.GetPokemon)
	return app
}
