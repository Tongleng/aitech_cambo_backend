package controllers

import (
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserController struct {
	Service services.UserService
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	users, err := c.Service.GetUsers()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return ctx.JSON(users)
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Parsing failed: " + err.Error()})
	}

	if user.Email == "" || user.Password == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error":             "Email and Password are required",
			"received_email":    user.Email,
			"received_password": user.Password,
		})
	}

	registeredUser, err := c.Service.Register(&user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Registration failed: email might already be in use",
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "Registration successful",
		"user": fiber.Map{
			"id":        registeredUser.ID,
			"email":     registeredUser.Email,
			"firstName": registeredUser.FirstName,
			"lastName":  registeredUser.LastName,
		},
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, token, err := c.Service.Login(req.Email, req.Password)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"user": fiber.Map{
			"id":        user.ID,
			"email":     user.Email,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"phone":     user.Phone,
		},
	})
}
