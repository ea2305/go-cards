package controllers

import (
	"encoding/json"
	"net/http"
)

var health_check = []string{"health_check"}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(health_check)
}
