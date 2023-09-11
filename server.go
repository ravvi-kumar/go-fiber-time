package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	lastTime := time.Now()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("last time " + lastTime.Format(time.UnixDate))
	})
	app.Get("/time", func(c *fiber.Ctx) error {
		lastTime = time.Now()
		fmt.Println("time", lastTime)
		return c.SendString("time updated " + lastTime.Format(time.UnixDate))
	})

	app.Listen(":3000")
}
