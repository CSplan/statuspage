package api

import (
	"encoding/json"
)

// Incident - A Statuspage incident creation request
type Incident struct {
	ID           string            `json:"id,omitempty"`
	Name         string            `json:"name,omitempty"`
	Status       string            `json:"status,omitempty"`
	ComponentIDs []string          `json:"component_ids,omitempty"`
	Components   map[string]string `json:"components,omitempty"` // Map of component status changes
}

// CreateIncident - Create an incident
func CreateIncident(title, componentID string) (incidentID string, e error) {
	incident := Incident{
		Name:   title,
		Status: "investigating",
		ComponentIDs: []string{
			componentID},
		Components: map[string]string{
			componentID: "partial_outage"}}
	r, e := doRequest("POST", route("/pages/"+pageID+"/incidents"), map[string]Incident{
		"incident": incident}, nil, 201)
	if e != nil {
		return "", e
	}
	json.NewDecoder(r.Body).Decode(&incident)
	return incident.ID, nil
}

// ResolveIncident - Resolve an incident
func ResolveIncident(incidentID string, componentID string) (e error) {
	incident := Incident{
		Status: "resolved",
		Components: map[string]string{
			componentID: "operational"}}
	_, e = doRequest("PATCH", route("/pages/"+pageID+"/incidents/"+incidentID), map[string]Incident{
		"incident": incident}, nil, 200)
	return e
}
