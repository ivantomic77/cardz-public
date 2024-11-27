package handlers

import (
	"encoding/json"
	"net/http"
)

var health = healthResponse{
	Status: "UP",
}

type healthResponse struct {
	Status string `json:"status"`
}

func GetHealth(w http.ResponseWriter, _ *http.Request) {
	response, err := json.Marshal(health)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
