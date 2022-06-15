package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[response-json] marshal exception \n")
		responseError(w, http.StatusInternalServerError, "response error")
		return
	}
	w.Write(json)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	var status = map[string]string{"error": "not found"}
	log.Printf("route not found: %v \n", r.RequestURI)
	responseJson(w, http.StatusNotFound, status)
}
