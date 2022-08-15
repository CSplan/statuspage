package api

// IncidentReq - A Statuspage incident creation request
type IncidentReq struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Status       string   `json:"status"`
	ComponentIDs []string `json:"component_ids"`
}

func NewIncident() {

}
