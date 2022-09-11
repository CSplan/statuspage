package main

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second}

func Monitor() {
	ticker := time.NewTicker(10 * time.Second) // Check status every 10 seconds
	for {
		components.load()
		incidents.load()
		components.API.Check("https://api.csplan.co", 404)
		components.Website.Check("https://csplan.co", 200)
		incidents.save()
		<-ticker.C
	}
}
