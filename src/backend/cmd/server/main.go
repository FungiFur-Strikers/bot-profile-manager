package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	openapi "bot-profile-manager/api"
	"bot-profile-manager/config"
	"bot-profile-manager/internal/api"
	"bot-profile-manager/internal/repository/mongodb"
	"bot-profile-manager/internal/service"
	mongodbclient "bot-profile-manager/pkg/mongodb"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client, err := mongodbclient.NewClient(cfg.MongoDBURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create a context that we'll use to gracefully shutdown our application
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	profileRepo := mongodb.NewProfileRepository(client, cfg.MongoDBName)
	profileService := service.NewProfileService(profileRepo)
	server := api.NewServer(profileService)

	r := gin.Default()
	openapi.RegisterHandlers(r, server)

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: r,
	}

	// Graceful server shutdown
	go func() {
		<-ctx.Done()
		log.Println("Shutting down server...")

		// Create a timeout context for shutdown
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("Server forced to shutdown: %v\n", err)
		}

		// Disconnect from MongoDB
		if err := client.Disconnect(shutdownCtx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v\n", err)
		}
	}()

	log.Printf("Server is running on %s\n", cfg.ServerAddress)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Server exited properly")
}
