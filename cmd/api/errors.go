package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"test/test/internal/response"
	"test/test/internal/validator"
)

func errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		logger.Error(err, debug.Stack())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func serverError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Error(err, debug.Stack())

	message := "The server encountered a problem and could not process your request"
	errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	errorMessage(w, r, http.StatusNotFound, message, nil)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	errorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func badRequest(w http.ResponseWriter, r *http.Request, err error) {
	errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func failedValidation(w http.ResponseWriter, r *http.Request, v validator.Validator) {
	err := response.JSON(w, http.StatusUnprocessableEntity, v)
	if err != nil {
		serverError(w, r, err)
	}
}
