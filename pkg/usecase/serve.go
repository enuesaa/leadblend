package usecase

import (
	"github.com/enuesaa/leadblend/pkg/controller"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(repos repository.Repos) error {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3001"},
	}))

	api := app.Group("api")
	api.Use(controller.HandleData)
	api.Use(controller.HandleError)
	api.GET("/spaces", controller.ListSpaces)
	// api.POST("/space/swicth")
	
	// api.GET("/planets")
	// api.GET("/planets/:planet_id")
	// api.GET("/planets/:planet_id/islands")
	// api.GET("/planets/:planet_id/islands/:island_id")
	// api.GET("/planets/:planet_id/islands/:island_id/tags")
	// api.GET("/planets/:planet_id/islands/:island_id/stones")

	// api.GET("/comets")

	app.Any("/*", ui.Serve)

	return app.Start(":3000")
}
