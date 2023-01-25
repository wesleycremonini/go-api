package main

import (
	"net/http"

	"test/test/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		serverError(w, r, err)
	}
}
