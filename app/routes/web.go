package routes

import (
	"GoFiberMVC/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// RegisterWebRoutes registers web routes
func RegisterWebRoutes(app *fiber.App) {
	userController := &controllers.UserController{}
	app.Get("", userController.Index)
}
