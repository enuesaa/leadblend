package main

import (
	"log"

	"github.com/enuesaa/leadblend/pkg/cli"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.New()

	app := &cobra.Command{
		Use:     "leadblend",
		Short:   "",
		Version: "0.0.1",
	}
	app.AddCommand(cli.CreateServeCmd(repos))
	app.AddCommand(cli.CreatePocCmd(repos))
	app.AddCommand(cli.CreateArchiveCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceErrors = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	if err := app.Execute(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
