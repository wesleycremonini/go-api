package main

import (
	"net/http"

	"github.com/alexedwards/flow"
)

func (app *application) routes() http.Handler {
	mux := flow.New()

	mux.NotFound = http.HandlerFunc(notFound)
	mux.MethodNotAllowed = http.HandlerFunc(methodNotAllowed)

	mux.HandleFunc("/status", app.status, http.MethodGet)

	{
		mux.HandleFunc("/wtfs", app.HandleWtfIndex, http.MethodGet)
		mux.HandleFunc("/wtfs/:id", app.HandleWtfFind, http.MethodGet)
	}

	return app.recoverPanic(mux)
}
