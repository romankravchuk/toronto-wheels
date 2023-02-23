package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/romankravchuk/logger"
	"github.com/romankravchuk/toronto-wheels/internal/server/router"
	loggermw "github.com/romankravchuk/toronto-wheels/internal/server/router/middleware/logger"
)

type Server struct {
	*fiber.App
}

func New(log *logger.Logger) *Server {
	s := &Server{App: fiber.New()}
	s.MountDefaultMiddlewares(log)
	s.Route("/api", func(r fiber.Router) {
		r.Mount("/v1", router.NewV1())
	})
	return s
}

func (s *Server) MountDefaultMiddlewares(log *logger.Logger) {
	s.Use(recover.New())
	s.Use(cors.New())
	s.Use(loggermw.New(loggermw.Config{Logger: log}))
}
