package main

import (
	"github.com/Lucasanim/shortly/config"
	"github.com/Lucasanim/shortly/internal/cache"
	"github.com/Lucasanim/shortly/internal/database"
	"github.com/Lucasanim/shortly/internal/migrations"
	"github.com/Lucasanim/shortly/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	cache.InitializeRedis()
	database.InitializeDb()
	migrations.Migrate()

	server := fiber.New()

	routes.Register(server)

	if err := server.Listen(":" + config.Env.Port); err != nil {
		panic(err)
	}
}
