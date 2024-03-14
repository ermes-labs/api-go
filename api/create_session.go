package api

import (
	"context"
)

// Commands to create a new session.
type CreateSessionCommands interface {
	// Creates a new session and returns the id of the session.
	CreateSession(
		ctx context.Context,
		opt CreateSessionOptions,
	) (string, error)
	// Returns the ids of the sessions.
	ScanSessions(
		ctx context.Context,
		cursor uint64,
		count int64,
	) (ids []string, newCursor uint64, err error)
}

// Create a new session, the options to define how the session is created. If no
// error is returned, a new session token is returned.
func (n *Node) CreateSession(
	ctx context.Context,
	opt CreateSessionOptions,
) (SessionToken, error) {
	sessionId, err := n.cmd.CreateSession(ctx, opt)

	// If there is an error, return it.
	if err != nil {
		return SessionToken{}, err
	}

	// Create a new session token.
	sessionLocation := NewSessionLocation(n.Host, sessionId)
	sessionToken := NewSessionToken(sessionLocation)

	// Return the session token.
	return sessionToken, nil
}

// Returns the ids of the sessions.
func (n *Node) ScanSessions(
	ctx context.Context,
	cursor uint64,
	count int64,
) (ids []string, newCursor uint64, err error) {
	return n.cmd.ScanSessions(ctx, cursor, count)
}
