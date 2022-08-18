package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CSplan/statuspage/api"
)

// Component - API component with polling state
type Component struct {
	*api.Component
	fails     []uint // failed poll timestamps, incident created at 6 within 1 minute
	successes []uint // successful poll timestamps (when an incident is active), incident resolved at 6 within 1 minute
}

// Components - API IDs of status components
type Components struct {
	API     *Component
	Website *Component
}

var components = Components{
	API:     &Component{},
	Website: &Component{}}

// Map Statuspage component IDs
func (c *Components) load() {
	list, err := api.GetComponents()
	if err != nil {
		panic(fmt.Errorf("Failed to get Statuspage components: %w", err))
	}

	// Map component IDs based on name
	for i, component := range list {
		switch component.Name {
		case "API":
			c.API.Component = &list[i]
		case "Website":
			c.Website.Component = &list[i]
		default:
			log.Printf("Unrecognized component: '%s' (status %s)", component.Name, component.Status)
		}
	}

	if c.API == nil {
		panic("API component not found")
	}
	if c.Website == nil {
		panic("Website component not found")
	}
}

func (c *Component) Check(url string, expectedStatus int) {
	r, err := client.Get(url)

	// Filter fail and success lists to only have timestamps within the past 60 seconds
	c.fails = filterTimestamps(c.fails)
	c.successes = filterTimestamps(c.successes)
	const pollTolerance = 6 // # of fails or successes within a minute required before an action is taken

	if err != nil || r.StatusCode != expectedStatus {
		c.successes = make([]uint, 0) // Reset successes
		c.fails = append(c.fails, uint(time.Now().Unix()))
		log.Printf("Failed check for %s (%d/%d)", c.Name, len(c.fails), pollTolerance)
		if len(c.fails) == pollTolerance && len(incidents[c.Name]) == 0 && c.Status == "operational" { // Only create an incident if no current incident exists for the API component
			log.Printf("Creating incident for %s", c.Name)
			incidentID, err := api.CreateIncident(fmt.Sprintf("%s outage", c.Name), c.ID)
			if err != nil {
				log.Println("Error creating incident:", err)
				return
			}
			// Store incident ID so it can be automatically resolved when the service is restored
			incidents[c.Name] = incidentID
			c.fails = make([]uint, 0)
		}
	} else if len(incidents[c.Name]) > 0 { // Resolve incidents that were automatically created
		c.fails = make([]uint, 0) // Reset fails
		c.successes = append(c.successes, uint(time.Now().Unix()))
		log.Printf("Successful check for %s (%d/%d)", c.Name, len(c.successes), pollTolerance)
		if len(c.successes) == pollTolerance {
			log.Printf("Resolving incident ID %s for %s", incidents[c.Name], c.Name)
			err = api.ResolveIncident(incidents[c.Name], c.ID)
			if err != nil {
				log.Println("Error resolving incident:", err)
				return
			}
			incidents[c.Name] = ""
			c.successes = make([]uint, 0)
		}
	}
}

// Filter a list of timestamps to only contain values within the past minute
func filterTimestamps(l []uint) []uint {
	now := uint(time.Now().Unix())
	for i, t := range l {
		if now-t > 60 {
			log.Println("Removing check older than 60 seconds")
			l = append(l[:i], l[i+1:]...)
		}
	}
	return l
}
