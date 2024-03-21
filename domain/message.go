package domain

// Message is a struct that represents a notification message
// the platform will send over the desired channel
type Message struct {
	Target    string
	Body      string
	Channel   string
	TrackingID string
}

