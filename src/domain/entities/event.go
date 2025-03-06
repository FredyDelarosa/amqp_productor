package entities

type Event struct {
	ID         int    `json:"id"`
	Zone       string `json:"zone"`
	DetectedAt string `json:"detected_at"`
}
