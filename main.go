package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go"
)

func main() {
	var app = fiber.New()
	var envErr = godotenv.Load(".env")

	app.Use(cors.New())

	if envErr != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
	}

	pusherClient := pusher.Client{
		AppID:   os.Getenv("APP_ID"),
		Key:     os.Getenv("KEY"),
		Secret:  os.Getenv("SECRET"),
		Cluster: os.Getenv("CLUSTER"),
		Secure:  true,
	}

	app.Post("/api/message", func(c *fiber.Ctx) error {
		var data map[string]string
		var err = c.BodyParser(&data)

		if err != nil {
			return c.JSON(err)
		}

		pusherClient.Trigger("chat", "message", data)
		return c.JSON(data)
	})

	app.Listen(":5000")
}
