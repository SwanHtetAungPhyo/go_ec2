package handler

import (
	"os"
	"time"

	"github.com/SwanHtetAungPhyo/auth/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	LoginUser := new(models.User)
	err := c.BodyParser(&LoginUser)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := &models.User{
		Email:    "swanhtet102002@gmail.com",
		Password: "Swanhtet12@",
	}

	if LoginUser.Email != user.Email || LoginUser.Password != user.Password {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	claims := jwt.MapClaims{
		"email": LoginUser.Email,
		"exp":   time.Now().Add(72 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "default_secret_key"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tok, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"access":  tok,
	})
}
