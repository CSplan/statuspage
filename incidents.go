package main

import (
	"encoding/json"
	"log"
	"os"
)

// Incidents - Open, automatically reported realtime incidents
type Incidents map[string]string

var incidents Incidents

func (i Incidents) filePath() string {
	return "incidents.json"
}

func (i *Incidents) load() {
	file, err := os.ReadFile(i.filePath())
	if err != nil {
		log.Println("Failed to read incidents.json:", err)
		return
	}

	json.Unmarshal(file, i)
}

func (i Incidents) save() {
	file, err := os.Create(i.filePath())
	if err != nil {
		log.Println("Failed to create incidents.json:", err)
		return
	}
	json.NewEncoder(file).Encode(i)
}
