package handlers

import (
	"fmt"
	"nanobgr-api/internal/queue"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func StatusHandler(r *queue.RedisClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing ID"})
		}

		statusKey := fmt.Sprintf("status:%s", id)
		statusStr, err := r.GetStatus(c.Context(), statusKey)

		if err == redis.Nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve status"})
		}

		response := fiber.Map{
			"id":     id,
			"status": statusStr,
		}

		if statusStr == "done" {
			response["result_image"] = fmt.Sprintf("processed/%s.png", id)
		}

		return c.JSON(response)
	}
}
