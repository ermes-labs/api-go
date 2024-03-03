package api

import (
	"encoding/json"
)

// SessionToken represents a session token.
type SessionToken struct {
	// The sessionLocation of the session.
	SessionLocation
}

// NewSessionToken creates a new session token.
func NewSessionToken(sessionLocation SessionLocation) SessionToken {
	return SessionToken{sessionLocation}
}

// Unmarshall a session token.
func UnmarshallSessionToken(sessionTokenBytes []byte) (*SessionToken, error) {
	if len(sessionTokenBytes) == 0 {
		return nil, nil
	}

	var sessionToken SessionToken
	err := json.Unmarshal(sessionTokenBytes, &sessionToken)

	if err != nil {
		return nil, err
	}

	return &sessionToken, nil
}

// Marshall a session token.
func MarshallSessionToken(sessionToken SessionToken) ([]byte, error) {
	sessionTokenBytes, err := json.Marshal(sessionToken)

	if err != nil {
		return nil, err
	}

	return sessionTokenBytes, nil
}
