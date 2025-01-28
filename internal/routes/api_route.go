package routes

import (
	"github.com/Lucasanim/shortly/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Register(server *fiber.App) {
	register := server.Group("/app")
	register.Post("/register", handlers.RegisterHandler)

	redirection := server.Group("/r")
	redirection.Get("/:hash", handlers.RedirectionHandler)
}
