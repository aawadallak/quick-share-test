package controller

import (
	"we/dto"
	"we/infra/repository"
	"we/usecases"

	"github.com/gofiber/fiber/v2"
)

type FileManagerController struct {
}

func (f *FileManagerController) UploadFile(c *fiber.Ctx) error {

	files, err := c.Context().MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dto := dto.FileDTO{}

	for _, file := range files.File {

		for range file {
			dto.File = append(dto.File, file...)
		}
	}

	svc := usecases.NewService(&repository.FileRepository{}, &repository.DatabaseRepository{})

	domain, err := svc.UploadFile(&dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": domain})

}

func (f *FileManagerController) DownloadFile(c *fiber.Ctx) error {

	id := c.Query("id", "empty")

	if id == "empty" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty id was provided",
		})
	}

	svc := usecases.NewService(&repository.FileRepository{}, &repository.DatabaseRepository{})

	path, err := svc.DownloadFile(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Download(*path)
}
