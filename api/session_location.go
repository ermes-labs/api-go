package api

// SessionLocation represents the sessionLocation of a session.
type SessionLocation struct {
	// The host.
	Host string `json:"host"`
	// The id of the session.
	SessionId string `json:"sessionId"`
}

// NewSessionLocation creates a new session sessionLocation.
func NewSessionLocation(host string, sessionId string) SessionLocation {
	return SessionLocation{
		Host:      host,
		SessionId: sessionId,
	}
}
