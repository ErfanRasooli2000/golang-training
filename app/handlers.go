package app

import (
	"encoding/json"
	"net/http"
)

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {

	writer.Header().Add("content-type", "application/json")

	writer.WriteHeader(code)

	json.NewEncoder(writer).Encode(data)
}
