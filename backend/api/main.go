package main

import (
	"server/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title Fiber Golang Rest Api
// @version 1.0
// @description this is swagger docs fpr rest api in golang fiber
// @host localhost:5000
// @BasePath /
// @schemes http
// securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @description Type "bearer" followed by a space  and token
func main() {

	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginsFunc: func(origin string) bool {
			return true

		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome")
	})
	//Server Swagger Documentation
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Listen(":5000")
}
