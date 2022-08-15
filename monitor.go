package main

import (
	"fmt"
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
		checkComponent(components.API, "https://api.csplan.co", 404)
		checkComponent(components.Website, "https://csplan.co", 200)
		incidents.save()
		<-ticker.C
	}
}

func checkComponent(c *api.Component, url string, expectedStatus int) {
	r, err := client.Get(url)
	if err != nil {
		log.Printf("Error getting %s status: %s", c.Name, err)
		return
	}

	if r.StatusCode != expectedStatus {
		if len(incidents[c.Name]) == 0 && c.Status == "operational" { // Only create an incident if no current incident exists for the API component
			incidentID, err := api.CreateIncident(fmt.Sprintf("%s outage", c.Name), c.ID)
			if err != nil {
				log.Println("Error creating incident:", err)
				return
			}
			// Store incident ID so it can be automatically resolved when the service is restored
			incidents[c.Name] = incidentID
		}
	} else if len(incidents[c.Name]) > 0 { // Resolve incidents that were automatically created
		err = api.ResolveIncident(incidents[c.Name], c.ID)
		if err != nil {
			log.Println("Error resolving incident:", err)
			return
		}
		incidents[c.Name] = ""
	}
}
