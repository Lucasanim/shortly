package handlers

import (
	"strings"

	"github.com/Lucasanim/shortly/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func RedirectionHandler(c *fiber.Ctx) error {
	hash := utils.CopyString(c.Params("hash"))

	url, err := services.LinkServiceImpl.GetUrl(hash)

	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	return c.Status(fiber.StatusMovedPermanently).Redirect(url)
}
