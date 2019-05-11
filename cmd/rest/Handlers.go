package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Status struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSha string `json:"lastcommitsha"`
}

type Statuses []Status

// RootHandler handles the call for the root / endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

// StatusHandler handles the call for the /status endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := Statuses{
		Status{Version: "v1.0", Description: "Some desp", LastCommitSha: "sha"},
	}

	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}
}
