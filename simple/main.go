package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Ping struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/", ping)

	port := "8000"
	fmt.Println("Ping API server run on port " + port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method == "GET" {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(&Ping{
			Status: "pong",
		})

		return
	}

	w.WriteHeader(http.StatusNotFound)
}
