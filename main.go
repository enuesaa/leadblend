package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/usecase"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.New()

	app := &cobra.Command{
		Use:     "leadblend",
		Short:   "",
		Version: "0.0.2",
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, _ := cmd.Flags().GetString("file")

			if !strings.HasSuffix(filename, ".leadblend") {
				return fmt.Errorf("invalid filename.")
			}
			repos.Using = filename
			if err := usecase.UsePkg(repos, filename); err != nil {
				return err
			}
			return usecase.Serve(repos)
		},
	}
	app.Flags().String("file", "main.leadblend", "filename to open")

	go func() {
		cancelch := make(chan os.Signal, 1)
		signal.Notify(cancelch, syscall.SIGTERM, syscall.SIGINT)
		<-cancelch
		if err := usecase.ClosePkg(repos, "main.leadblend"); err != nil {
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
