package server

import (
	"log"
	"os"
	"quick_share/server/middlewares"
	"quick_share/server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Port   string
	Server *fiber.App
}

func NewServer() Server {
	return Server{
		Port: os.Getenv("SERVER_PORT"),
		Server: fiber.New(fiber.Config{
			BodyLimit: 100 * 1024 * 1024,
		}),
	}
}

func (s *Server) Run() {

	s.Server.Use(logger.New(logger.ConfigDefault))

	s.Server.Use(middlewares.CorsMiddleware())

	router := routes.SetRoute(s.Server)
	log.Fatalln(router.Listen(s.Port))

}
