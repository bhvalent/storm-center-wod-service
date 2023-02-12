package main

import (
	"fmt"
	"log"
	"os"

	"storm-center-wod-service/internal/app/routes"
	"storm-center-wod-service/internal/domain/models"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load ennvironment variables: %e", err)
	}

	var cfg models.Config
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Unable to parse ennvironment variables: %e", err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := models.Application{
		Logger: logger,
		Config: cfg,
	}

	router := routes.Routes(&app)

	err = router.Run(fmt.Sprintf("%s:%d", cfg.BaseUrl, cfg.Port))
	if err != nil {
		log.Println(err)
	}
}
