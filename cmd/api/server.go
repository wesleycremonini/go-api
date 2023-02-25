package main

import (
	"errors"
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/caddyserver/certmagic"
)

const (
	defaultIdleTimeout  = time.Minute
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

func (app *application) run() error {
	var err error
	if app.config.domainName == "localhost" {
		srv := &http.Server{
			Addr:         ":80",
			Handler:      app.routes(),
			ErrorLog:     log.New(logger, "", 0),
			IdleTimeout:  defaultIdleTimeout,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		}
		logger.Info("starting dev server")
		err = srv.ListenAndServe()
	} else {
		fmt.Println("starting certmagic server")
		certmagic.DefaultACME.Agreed = true
		certmagic.DefaultACME.Email = app.config.email
		certmagic.DefaultACME.CA = certmagic.LetsEncryptProductionCA
		certmagic.Default.Storage = &certmagic.FileStorage{Path: "/certs"}
		certmagic.HTTPPort, certmagic.HTTPSPort = 80, 443
		err = certmagic.HTTPS([]string{app.config.domainName, "www." + app.config.domainName}, app.routes())
	}
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	logger.Info("server stopped")
	app.wg.Wait()
	return nil
}
