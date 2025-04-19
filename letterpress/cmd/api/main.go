package main

import (
	"os"

	"letterpress/db"

	"letterpress/handler"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {

	var err error
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	dbInstance, err := db.Init(logger)
	if err != nil {
		logger.Err(err).Msg("Connection failed")
		os.Exit(1)
	}
	logger.Info().Msg("Database connection established")

	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		logger.Err(err).Msg("Connection failed")
		os.Exit(1)
	}

	h := handler.New(dbInstance, esClient, logger)
	router := gin.Default()
	rg := router.Group("/v1")
	h.Register(rg)
	router.Run(":8080")
}
