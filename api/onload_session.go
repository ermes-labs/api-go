package api

import (
	"context"
	"io"
)

// Commands to onload a session.
type OnloadSessionCommands interface {
	// StartOnload starts the onload of a session and returns the id of the
	// session.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	// - ErrSessionAlreadyOnloaded: If the session is already onloaded.
	OnloadSession(
		ctx context.Context,
		metadata SessionMetadata,
		reader io.Reader,
		opt OnloadSessionOptions,
	) (string, error)
}

// Setup the onload of a session. The function returns the location of the
// session, a function to onload the session data, a rollback function to defer
// that deletes the session if the onload fails, and an error.
// errors:
// - ErrSessionAlreadyOnloaded: If the session is already onloaded.
func (n *Node) OnloadSession(
	cmd OnloadSessionCommands,
	ctx context.Context,
	metadata SessionMetadata,
	reader io.Reader,
	opt OnloadSessionOptions,
) (SessionLocation, error) {
	// Start the onload of the session.
	sessionId, err := cmd.OnloadSession(ctx, metadata, reader, opt)

	// If there is an error, return it.
	if err != nil {
		return SessionLocation{}, err
	}

	// Return the location of the session.
	return NewSessionLocation(n.Host, sessionId), nil
}
