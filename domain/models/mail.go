package models

// Mail is simplify mail json object
type Mail struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body,string"`
	Html    string   `json:"html,string"`
}
