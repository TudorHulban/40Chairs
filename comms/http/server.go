package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	engine *fiber.App
}

func NewServer() (*Server, error) {
	s := Server{
		engine: fiber.New(),
	}

	s.addRoutes()

	return &s, nil
}

func (s *Server) addRoutes() {
	s.engine.Get("/"+urlPing, s.ping)
}

func (s *Server) ping(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
