package handlers

import (
	"encoding/json"
	"fmt"
	"nanobgr-api/internal/config"
	"nanobgr-api/internal/queue"
	"nanobgr-api/internal/storage"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ValidationResponse struct {
	Error string `json:"error"`
}

type TaskPayload struct {
	ID       string `json:"id"`
	RawImage string `json:"raw_image"`
}

var allowedTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
}

func UploadHandler(cfg *config.Config, s *storage.MinioClient, r *queue.RedisClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fileHeader, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(ValidationResponse{Error: "Missing image file"})
		}

		// Generate ID
		id := uuid.New().String()
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if ext == "" {
			ext = ".jpg" // Default fallback
		}
		rawObjectName := fmt.Sprintf("raw/%s%s", id, ext)

		// Open file
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ValidationResponse{Error: "Failed to open file"})
		}
		defer file.Close()

		// Upload to MinIO
		err = s.UploadFile(c.Context(), rawObjectName, file, fileHeader.Size, fileHeader.Header.Get("Content-Type"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ValidationResponse{Error: "Failed to upload to storage: " + err.Error()})
		}

		// Write initial status to Redis
		statusKey := fmt.Sprintf("status:%s", id)
		err = r.SetStatus(c.Context(), statusKey, "processing")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ValidationResponse{Error: "Failed to initialize state"})
		}

		// Push to Queue
		payload := TaskPayload{
			ID:       id,
			RawImage: rawObjectName,
		}
		payloadBytes, _ := json.Marshal(payload)
		err = r.PushTask(c.Context(), "image_queue", string(payloadBytes))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ValidationResponse{Error: "Failed to queue task"})
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"id":     id,
			"status": "processing",
		})
	}
}
