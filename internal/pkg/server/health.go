package server

import (
	"log"
	"net/http"
)


func ServeHealthCheck(healthPath string, healthAddress string) {
	http.HandleFunc("/" + healthPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([] byte(`{"status": "OK"}`));  err != nil {
			log.Fatalf("Error write health check: %s", err.Error())
		}
	})
	
	if err := http.ListenAndServe(healthAddress, nil); err != nil {
		log.Fatalf("Error serving health check endpoint: %s", err.Error())
	}
}