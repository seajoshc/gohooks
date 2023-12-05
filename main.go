package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Structs to match Compass component JSON structure
type Component struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

type Parameters map[string]string

type WebhookPayload struct {
	Component  Component  `json:"component"`
	Parameters Parameters `json:"parameters"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Process the payload here
	fmt.Printf("Received component: %+v\n", payload.Component)
	fmt.Printf("-->with parameters: %+v\n", payload.Parameters)

	// Respond to the webhook sender
	fmt.Fprintf(w, "Webhook received and processed")
}

func main() {
	port := ":8888"
	fmt.Printf("Starting webhooks server on localhost:%s/compass\n", port)
	fmt.Println("ngrok http 8888")
	http.HandleFunc("/compass", webhookHandler)
	log.Panic(http.ListenAndServe(port, nil))
}
