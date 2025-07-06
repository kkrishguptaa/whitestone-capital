package main

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/kkrishguptaa/whitestone-capital/apps/api/lib"
)

func main() {
	lib.InitEnv()
	lib.InitDatabase()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(requestid.New())
	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(helmet.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(lib.Env)
	})

	log.Fatal(app.Listen(":" + lib.Env.API_PORT))
}
