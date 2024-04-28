package cli

import (
	"github.com/enuesaa/leadblend/pkg/controller"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
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

			app.Any("/*", ui.Serve)

			return app.Start(":3000")
		},
	}

	return cmd
}
