package handlers

import (
	"encoding/json"
	"net/http"

	"platzi.com/go/rest-ws-go/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //lo que le respondo al cliente es un json
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to Platzi GO",
			Status:  true,
		})
	}
}
