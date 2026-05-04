package main

import (
    "github.com/gofiber/fiber/v2"
    "log"
)

func main() {
    app := fiber.New()

    // Health check route
    app.Get("/api/health", func(c *fiber.Ctx) error {
        return c.Status(200).JSON(fiber.Map{
            "status": "success",
            "message": "Go Fiber Layered MVC is live",
        })
    })

    log.Fatal(app.Listen(":8080")) // Standard port for Go backends
}