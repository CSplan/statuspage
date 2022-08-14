package api

import (
	"encoding/json"
)

type Component struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func GetComponents() (components []Component, e error) {
	// Page query params will be required by Statuspage in Q1 2023
	const query = "?page=1&per_page=100"

	r, e := doRequest("GET", route("/pages/"+pageID+"/components"+query), nil, nil, 200)
	if e != nil {
		return nil, e
	}

	// Decode body
	json.NewDecoder(r.Body).Decode(&components)

	return components, nil
}
