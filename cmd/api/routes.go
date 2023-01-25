package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(notFound)
	mux.MethodNotAllowed = http.HandlerFunc(methodNotAllowed)

	mux.HandlerFunc(http.MethodGet, "/status", app.status)

	{
		mux.HandlerFunc(http.MethodGet, "/wtfs", app.HandleWtfIndex)
		mux.HandlerFunc(http.MethodGet, "/wtfs/:id", app.HandleWtfFind)
	}

	return app.recoverPanic(mux)
}
