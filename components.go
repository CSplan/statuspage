package main

import (
	"fmt"
	"log"

	"github.com/CSplan/statuspage/api"
)

// Components - API IDs of status components
type Components struct {
	API     string
	Website string
}

var components Components

// Map Statuspage component IDs
func loadComponents() {
	list, err := api.GetComponents()
	if err != nil {
		panic(fmt.Errorf("Failed to get Statuspage components: %w", err))
	}

	// Map component IDs based on name
	for _, c := range list {
		switch c.Name {
		case "API":
			components.API = c.ID
		case "Website":
			components.Website = c.ID
		default:
			log.Printf("Unrecognized component: '%s' (status %s)", c.Name, c.Status)
		}
	}

	if len(components.API) == 0 {
		panic("API component not found")
	}
	if len(components.Website) == 0 {
		panic("Website component not found")
	}
}
