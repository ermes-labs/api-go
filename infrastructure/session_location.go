package infrastructure

// SessionLocation represents the sessionLocation of a session.
type SessionLocation struct {
	// The host.
	Host string
	// The id of the session.
	SessionId string
}

// NewSessionLocation creates a new session sessionLocation.
func NewSessionLocation(host string, sessionId string) SessionLocation {
	return SessionLocation{
		Host:      host,
		SessionId: sessionId,
	}
}
