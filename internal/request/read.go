package request

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alexedwards/flow"
)

func ReadIDParam(r *http.Request) (int, error) {
	id, err := strconv.ParseInt(flow.Param(r.Context(), "id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return int(id), nil
}
