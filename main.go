package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enuesaa/leadblend/pkg/cli"
	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/usecase"
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

    go func() {
		cancelch := make(chan os.Signal, 1)
		signal.Notify(cancelch, syscall.SIGTERM, syscall.SIGINT)
		<-cancelch
		if err := usecase.ClosePkg(repos, "main"); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
        os.Exit(0)
    }()

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
