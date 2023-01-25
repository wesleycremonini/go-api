package main

import (
	"encoding/json"
	"net/http"
	"test/test/internal/domain"
	"test/test/internal/request"
	"test/test/internal/response"

	"github.com/julienschmidt/httprouter"
)

func (app *application) HandleWtfIndex(w http.ResponseWriter, r *http.Request) {
	var filter domain.WtfFilter

	params := httprouter.ParamsFromContext(r.Context())
	idsString := params.ByName("ids")
	wtfsString := params.ByName("wtfs")

	var idsInput []int
	if idsString != "" {
		if err := json.Unmarshal([]byte(idsString), &idsString); err != nil {
			panic(err)
		}
	}
	
	var wtfsInput []string
	if wtfsString != "" {
		if err := json.Unmarshal([]byte(wtfsString), &wtfsString); err != nil {
			panic(err)
		}
	}


	filter = domain.WtfFilter{
		IDs:  idsInput,
		Wtfs: wtfsInput,
	}

	wtfs, n, err := app.WtfService.Wtfs(filter)
	if err != nil {
		serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, response.Envelope{"wtfs": wtfs, "results_qty": n})
	if err != nil {
		serverError(w, r, err)
		return
	}
}

func (app *application) HandleWtfFind(w http.ResponseWriter, r *http.Request) {
	id, err := request.ReadIDParam(r)
	if err != nil {
		badRequest(w, r, err)
		return
	}

	wtf, err := app.WtfService.WtfByID(id)
	if err != nil {
		notFound(w, r)
		return
	}

	response.JSON(w, http.StatusOK, response.Envelope{"wtf": wtf})
}
