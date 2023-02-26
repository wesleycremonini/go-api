package main

import (
	"errors"

	"log"
	"net/http"
	"time"
)

const (
	defaultIdleTimeout  = time.Minute
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

func (app *application) run() error {
	var err error
	srv := &http.Server{
		Addr:         ":80",
		Handler:      app.routes(),
		ErrorLog:     log.New(logger, "", 0),
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}
	logger.Info("starting server")
	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	logger.Info("server stopped")
	app.wg.Wait()
	return nil
}
