package main

import (
	"log"
	"net/http"
	"time"

	"github.com/CSplan/statuspage/api"
)

var client = &http.Client{}

func monitor() {
	ticker := time.NewTicker(10 * time.Second) // Check status every 10 seconds
	for {
		components.load()
		incidents.load()
		checkComponent(components.API, "https://api.csplan.co")
		<-ticker.C
	}
}

func checkComponent(c *api.Component, url string) {
	r, err := client.Get(url)
	if err != nil {
		log.Println("Error getting API status:", err)
		return
	}
	status := r.StatusCode

	if status%500 < 100 { // Create realtime incidents for 5xx responses
		if c.Status == "operational" { // Only create an incident if no current incident exists for the API component
			// Create incident
		}
	} else if len(incidents[c.Name]) > 0 { // Resolve incidents that were automatically created
		// Resolve incident
	}
}
