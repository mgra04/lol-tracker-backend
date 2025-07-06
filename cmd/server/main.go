package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/mgra04/backend-lol-v1/internal/api/handler"
	"github.com/mgra04/backend-lol-v1/internal/cache"
	"github.com/mgra04/backend-lol-v1/internal/config"
	"github.com/mgra04/backend-lol-v1/internal/constants"
	"github.com/mgra04/backend-lol-v1/internal/logger"
	"github.com/mgra04/backend-lol-v1/internal/repository/mongo"
	"github.com/mgra04/backend-lol-v1/internal/repository/riotapi"
	"github.com/mgra04/backend-lol-v1/internal/service"
)

func main() {
	// Logger init
	logger.SetupLogger()

	// Config load
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config.")
	}

	// MongoDB init
	// mongoClient, err := mongo.NewMongoClient(cfg.MongoURI)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to connect to MongoDB.")
	// }
	// defer mongoClient.Disconnect(context.Background())
	// MongoDB init
	mongoClient, err := mongo.NewMongoClient(cfg.MongoURI)
	if err != nil {
		log.Warn().Err(err).Msg("Could not connect to MongoDB. Continuing without DB.")
		mongoClient = nil // fallback: brak repo Mongo
	} else {
		defer mongoClient.Disconnect(context.Background())
	}
	
	// BigCache init
	cacheClient := cache.NewBigCache()

	// Riot API client init
	riotClient := riotapi.NewRiotApiClient(cfg.RiotAPIKey)

	// Service layer init
	profileService := service.NewProfileService(riotClient, mongoClient, cacheClient)

	// router init
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// routes
	r.Get("/api/profile", handler.NewProfileHandler(profileService).GetProfile)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
		ReadTimeout: constants.ReadTimeout,
		WriteTimeout:  constants.WriteTimeout,
		IdleTimeout:  constants.IdleTimeout,
	}

	// Starting server log
	log.Info().Msgf("Starting server on port %d.", cfg.Port)
	if err := srv.ListenAndServe(); err !=nil {
		log.Fatal().Err(err).Msg("Server crashed.")
	}
}