package controller

import (
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/pniewiarowski/simple-rest-api/env"
	"github.com/pniewiarowski/simple-rest-api/models"
)

func Register(ctx *fiber.Ctx) error {
	auth := new(models.Auth)

	if err := ctx.BodyParser(&auth); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": err})
	}

	if _, err := mail.ParseAddress(auth.Email); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": "incorrect email"})
	}

	if len(auth.Password) < 8 {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": "password should be minimum 8 characters"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"error": err})
	}

	auth.Password = string(hashedPassword)
	if err = models.CreateAuth(auth); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"message": "account created"})
}

func Login(ctx *fiber.Ctx) error {
	userAuth := new(models.Auth)

	if err := ctx.BodyParser(&userAuth); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": err})
	}

	auth := models.GetAuthByEmail(userAuth.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(userAuth.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"error": "incorrect password"})
	}

	if !auth.Enable {
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"message": "account is disable"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = auth.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenStringed, err := token.SignedString([]byte(env.GetPrivateKey()))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"error": err})
	}

	return ctx.JSON(&fiber.Map{"token": tokenStringed})
}
