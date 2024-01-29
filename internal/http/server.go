package http

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/fiber/v2/middleware/recover"

type Server struct {
	FiberApp              *fiber.App
	mockPartnerController *PartnerApiController
	interactionController *MockController
}

func NewServer(mockPartnerController *PartnerApiController, controller *MockController) *Server {
	srv := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	recoverConfig := recover.ConfigDefault
	recoverConfig.EnableStackTrace = true
	srv.Use(recover.New(recoverConfig))

	return &Server{
		FiberApp:              srv,
		mockPartnerController: mockPartnerController,
		interactionController: controller,
	}
}
