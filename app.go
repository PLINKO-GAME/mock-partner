package main

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/http"
	"bitbucket.org/1-pixel-games/mock-partner/internal/partner"
	log "github.com/sirupsen/logrus"
)

type application struct {
	server *http.Server
}

func newApplication(config *config) (*application, error) {
	var app = new(application)

	greeting(config)

	partnerService := partner.New(config.CoreURL)
	mockPartnerController := http.NewPartnerApiController(partnerService)
	interactionController := http.NewMockController(partnerService)
	app.server = http.NewServer(mockPartnerController, interactionController)
	app.server.WithPartnerApiRoutes()
	app.server.WithMockRoutes()
	app.server.WithHealth()

	return app, nil
}

func greeting(config *config) {
	log.Infof("Listening on port %s", config.HTTPPort)
}
