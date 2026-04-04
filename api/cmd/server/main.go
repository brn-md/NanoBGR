package main

import (
	"fmt"
	"log"

	"nanobgr-api/internal/config"
	"nanobgr-api/internal/handlers"
	"nanobgr-api/internal/queue"
	"nanobgr-api/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize MinIO
	minioClient, err := storage.NewMinioClient(
		cfg.MinioEndpoint,
		cfg.MinioAccessKey,
		cfg.MinioSecretKey,
		cfg.MinioBucket,
		cfg.MinioUseSSL,
	)
	if err != nil {
		log.Fatalf("Failed to initialize MinIO: %v", err)
	}

	// Initialize Redis
	redisClient, err := queue.NewRedisClient(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	// Setup Fiber app
	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 50MB max upload size
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // For development purposes; restrict this in production
		AllowMethods: "GET,POST,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/upload", handlers.UploadHandler(cfg, minioClient, redisClient))
	app.Get("/status/:id", handlers.StatusHandler(redisClient))

	log.Printf("Server listening on port %s", cfg.Port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.Port)))
}
