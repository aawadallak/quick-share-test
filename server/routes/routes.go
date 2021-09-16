package routes

import (
	"we/app/http/controller"

	"github.com/gofiber/fiber/v2"
)

func SetRoute(router *fiber.App) *fiber.App {

	controller := controller.FileManagerController{}

	router.Post("/upload", controller.UploadFile)

	router.Get("/download/", controller.DownloadFile)

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pong"})
	})

	return router
}
