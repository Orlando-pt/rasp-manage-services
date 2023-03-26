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

	app.Get("/:option/series", func(c *fiber.Ctx) error {
		if c.Params("option") == "rm" {
			cmd := exec.Command("rm", "-rf", "/home/rasp-server/media/series/Two_and_a_half_men/*")
			err := cmd.Run()
			return answerRequest(c, err)
		} else if c.Params("option") == "mv" {
			cmd := exec.Command("mv", "/home/qbittorrent/Downloads/*", "/home/rasp-server/media/series/Two_and_a_half_men/")
			err := cmd.Run()
			return answerRequest(c, err)
		} else {
			return c.SendString("Invalid option")
		}
	})

	app.Get("/df", func(c *fiber.Ctx) error {
		cmd := exec.Command("df", "-h")
		out, err := cmd.Output()

		if err != nil {
			return c.SendString(err.Error())
		}

		return c.SendString(string(out))
	})

	app.Listen(":3001")
}

func answerRequest(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString("Operation completed successfully")
}
