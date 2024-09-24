package healthcheck

import (
	responses "template-api-fiber-golang/src/server/response"

	"github.com/gofiber/fiber/v2"
)

// Register route to health check endpoint and implement its logical function response that must be {message:"OK"}.
func RegisterRoutes(router fiber.Router) {
	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(responses.OK())
	})
}
