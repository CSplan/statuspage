package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

// Incidents - Open, automatically reported realtime incidents
type Incidents map[string]string

var incidents = make(Incidents)

func (i Incidents) filePath() string {
	return "incidents.json"
}

func (i *Incidents) load() {
	file, err := os.Open(i.filePath())
	if errors.Is(err, os.ErrNotExist) {
		i.save()
		return
	} else if err != nil {
		log.Println("Failed to read incidents.json:", err)
		return
	}

	json.NewDecoder(file).Decode(i)
}

func (i Incidents) save() {
	file, err := os.Create(i.filePath())
	if err != nil {
		log.Println("Failed to create incidents.json:", err)
		return
	}
	json.NewEncoder(file).Encode(i)
}
