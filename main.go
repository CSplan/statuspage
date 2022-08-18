package main

import (
	"github.com/CSplan/statuspage/api"
)

func main() {
	// Init API config
	config := ParseConfig()
	api.SetKey(config.Key)
	api.SetPageID(config.PageID)

	// Begin monitor loop
	Monitor()
}
