package app

import "github.com/gofiber/fiber/v2"

// Route that stores a path, with params inside, the method type, the handler and any decorators
type Route struct {
	Path          string
	Method        Method
	Middleware    []fiber.Handler
	Handler       fiber.Handler
	StaticHandler string
}
