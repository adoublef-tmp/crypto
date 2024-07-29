package http

import "github.com/gofiber/fiber/v2"

func Handler(alerter Alerter) *fiber.App {
	app := fiber.New()
	app.Post("/alert", handleAlert(alerter))
	return app
}

func handleAlert(alerter Alerter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Params("email")
		err := alerter.Alert(c.Context(), email)
		if err != nil {
			return fiber.ErrInternalServerError
		}
		return c.SendString("alert")
	}
}
