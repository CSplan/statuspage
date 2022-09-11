package main

import (
	"net/http"
	"sync"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second}

func monitor() {
	ticker := time.NewTicker(10 * time.Second) // Check status every 10 seconds
	for {
		components.load()
		incidents.load()
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			components.API.Check("https://api.csplan.co", 404)
			wg.Done()
		}()
		go func() {
			components.Website.Check("https://csplan.co", 200)
			wg.Done()
		}()
		wg.Wait() // Wait for all checks to complete
		incidents.save()
		<-ticker.C
	}
}
