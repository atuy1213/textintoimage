package handler

import (
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := os.ReadFile("../src/sampale.jpeg")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}
