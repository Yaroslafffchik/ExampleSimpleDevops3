package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Info struct {
	Service   string `json:"service"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Host      string `json:"host"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	service := os.Getenv("SERVICE_NAME")
	if service == "" {
		service = "service2"
	}
	info := Info{
		Service:   service,
		Message:   "Greetings from " + service,
		Timestamp: time.Now().Format(time.RFC3339),
		Host:      r.Host,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/health", healthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting %s on :%s\n", os.Getenv("SERVICE_NAME"), port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
