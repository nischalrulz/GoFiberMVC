package controllers

import (

	"github.com/gofiber/fiber/v2"
)

// UserController handles user-related requests
type UserController struct {
	
}

// Index handles the GET request for listing users
func (c *UserController) Index(ctx *fiber.Ctx) error {
	// Render index template
	return ctx.Render("index", fiber.Map{
		"Title": "Error",
	})
}

