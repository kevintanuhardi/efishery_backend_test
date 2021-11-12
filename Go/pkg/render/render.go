package render

import (
	"encoding/json"
	"net/http"
)

type empty struct{}

var EmptyResponse empty

func Response(w http.ResponseWriter, httpStatusHeader int, data, errors, meta interface{}) {
	apiResponse := struct {
		Data   interface{} `json:"data"`
		Errors interface{} `json:"errors"`
		Meta   interface{} `json:"meta"`
	}{
		data,
		errors,
		meta,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusHeader)

	_ = json.NewEncoder(w).Encode(apiResponse)
}
