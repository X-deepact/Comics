package route

import (
	"github.com/gofiber/fiber/v2"
)

func SetupFileRoute(rootFilePath string, router fiber.Router) {
	router.Static("/files", rootFilePath)
}
