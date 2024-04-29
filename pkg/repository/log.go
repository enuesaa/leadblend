package repository

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
}

type LogRepositoryInterface interface {
	Info(format string, v ...any)
	Fatal(err error)
}
type LogRepository struct{}

func (repo *LogRepository) Info(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	log.Printf("%s\n", message)
}

func (repo *LogRepository) Fatal(err error) {
	log.Fatalf("Error: %s\n", err.Error())
}
