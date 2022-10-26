package common

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON common response
func ResponseJSON(w http.ResponseWriter, statusCode int, respStruct interface{}) int {
	b, err := json.Marshal(&respStruct)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return 0
	}
	w.WriteHeader(statusCode)
	write, err := w.Write(b)
	if err != nil {
		return 0
	}
	return write
}
