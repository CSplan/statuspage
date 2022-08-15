package main

import (
	"fmt"
	"log"

	"github.com/CSplan/statuspage/api"
)

// Components - API IDs of status components
type Components struct {
	API     *api.Component
	Website *api.Component
}

var components Components

// Map Statuspage component IDs
func (c *Components) load() {
	list, err := api.GetComponents()
	if err != nil {
		panic(fmt.Errorf("Failed to get Statuspage components: %w", err))
	}

	// Map component IDs based on name
	for _, component := range list {
		switch component.Name {
		case "API":
			c.API = &component
		case "Website":
			c.Website = &component
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
