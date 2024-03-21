package postgres

import "time"

// Client is a struct that holds a connection to a postgres database
type Client struct {
}

// NewClient returns a new instance of the client
func NewClient() *Client {
	return &Client{}
}

// BatchBlackList will mark a list of emails as blacklisted
func (c *Client) BatchBlackList(email []string) error {
	// simulate a long running process
	time.Sleep(500 * time.Millisecond)
	return nil
}
