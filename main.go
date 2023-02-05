package main

import (
	"github.com/gofiber/fiber/v2"
	"os/exec"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/:option/jellyfin", func(c *fiber.Ctx) error {
		if c.Params("option") == "start" {
			cmd := exec.Command("systemctl", "restart", "jellyfin.service")
			err := cmd.Run()
			return answerRequest(c, err)
		} else if c.Params("option") == "stop" {
			cmd := exec.Command("systemctl", "stop", "jellyfin.service")
			err := cmd.Run()
			return answerRequest(c, err)
		} else {
			return c.SendString("Invalid option")
		}
	})

	app.Get("/:option/torrent", func(c *fiber.Ctx) error {
		if c.Params("option") == "start" {
			cmd := exec.Command("systemctl", "restart", "qbittorrent.service")
			err := cmd.Run()
			return answerRequest(c, err)
		} else if c.Params("option") == "stop" {
			cmd := exec.Command("systemctl", "stop", "qbittorrent.service")
			err := cmd.Run()
			return answerRequest(c, err)
		} else {
			return c.SendString("Invalid option")
		}
	})

	app.Listen(":3001")
}

func answerRequest(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString("Operation completed successfully")
}
