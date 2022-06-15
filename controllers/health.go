package controllers

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var status = map[string]string{"status": "ok"}
	responseJson(w, http.StatusOK, status)
}
