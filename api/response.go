package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

// sendError sending response with error based on error + status code.
func sendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := http.StatusInternalServerError

	w.WriteHeader(statusCode)

	err = json.NewEncoder(w).Encode(Error{Error: err.Error()})
	if err != nil {
		log.Println(err)
	}
}
