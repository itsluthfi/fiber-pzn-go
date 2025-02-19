package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	// prefork
	if fiber.IsChild() {
		fmt.Println("This is child process")
	} else {
		fmt.Println("This is parent process")
	}

	// middleware
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("I'm middleware before processing request")
		err := c.Next()
		fmt.Println("I'm middleware after processing request")
		return err
	})

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
