package main

import (
	"fmt"

	"github.com/CSplan/statuspage/api"
)

func main() {
	// Init API config
	config := ParseConfig()
	api.SetKey(config.Key)
	api.SetPageID(config.PageID)

	components, err := api.GetComponents()
	if err != nil {
		panic(err)
	}
	fmt.Println(components)
}
