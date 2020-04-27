package http_utils

import (
	"net/http"
	"encoding/json"
	"github.com/sampila/go-utils/rest_errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
  response, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, err rest_errors.RestErr) {
	RespondJson(w, err.Status(), err)
}
