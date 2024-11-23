package main

import (
	"github.com/SwanHtetAungPhyo/auth/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
		AllowHeaders: "Content-Type,application/json",
	}))
	// key := os.Getenv("SECRET_KEY")
	
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(key)},
	// }))
	app.Post("/user/login",handler.Login)
	app.Get("/user/info",func (c * fiber.Ctx) error  {
		return c.SendString("Hello from the go user service")
	})
	app.Listen(":8001")
	
}