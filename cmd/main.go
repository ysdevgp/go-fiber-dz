package main

import (
	"log"
	"strconv"
	"ysd/go-fiber-dz/config"
	"ysd/go-fiber-dz/internal/pages"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	config.Init()
	appConfig := config.GetAppConfig()

	cfg := fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			if fiberErr, ok := err.(*fiber.Error); ok && fiberErr.Code == fiber.StatusNotFound {
				return c.Status(fiber.StatusNotFound).SendString("HTTP 404 Not Found")
			}
			const errString string = "HTTP 500 Internal Server Error"
			return c.Status(fiber.StatusInternalServerError).SendString(errString)
		},
	}

	app := fiber.New(cfg)
	app.Use(recover.New())

	pages.NewHandler(app)

	if err := app.Listen(":" + strconv.Itoa(appConfig.Port)); err != nil {
		log.Fatal(err.Error())
	}
}
