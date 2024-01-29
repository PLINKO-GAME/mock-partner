package main

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	conf, err := newConfig()
	if err != nil {
		log.Fatalf("failed to apply configuration: %v", err)
	}

	app, err := newApplication(conf)
	if err != nil {
		log.WithError(err).Fatal("failed to bootstrap the app")
	}

	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, _ := errgroup.WithContext(appCtx)
	group.Go(
		func() error {
			return app.server.FiberApp.Listen(conf.HTTPPort)
		},
	)

	err = group.Wait()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.WithError(err).Fatal("context canceled")
		} else {
			log.WithError(err).Fatal("error received")
		}
	}
}
