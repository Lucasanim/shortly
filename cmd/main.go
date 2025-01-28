package main

import (
	"github.com/Lucasanim/shortly/config"
	"github.com/Lucasanim/shortly/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	routes.Register(server)

	server.Listen(config.Env.Port)
}
