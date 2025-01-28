package handlers

import (
	"github.com/Lucasanim/shortly/config"
	"github.com/Lucasanim/shortly/internal/models"
	"github.com/Lucasanim/shortly/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	var createLink models.CreateLink

	if err := c.BodyParser(&createLink); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	shortenedLink := services.LinkServiceImpl.Create(createLink)
	linkAddress := config.Env.BaseUrl + "/r/" + shortenedLink.Hash

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":        "Link created",
		"shortened link": linkAddress,
	})
}
