package api

import (
	"context"
)

// Commands to create and acquire a session.
type CreateAndAcquireSessionCommands interface {
	AcquireSessionCommands
	// Creates a new session and acquires it. Returns the id of the session.
	CreateAndAcquireSession(
		ctx context.Context,
		options CreateAndAcquireSessionOptions,
	) (string, error)
}

// Create a new session and acquire it, the run the ifCreatedAndAcquired
// callback. Inside the callback is possible to safely use the session. The
// options defines how the session is created and acquired.
//
// If the callback is run, a SessionToken is returned, otherwise an error is
// returned.
func (n *Node) CreateAndAcquireSession(
	ctx context.Context,
	opt CreateAndAcquireSessionOptions,
	ifCreatedAndAcquired func(sessionToken SessionToken) error,
) (SessionToken, error) {
	// Create and acquire the session.
	sessionId, err := n.cmd.CreateAndAcquireSession(ctx, opt)

	// If there is an error, return it.
	if err != nil {
		return SessionToken{}, err
	}

	// Defer the release of the session.
	defer func() {
		n.cmd.ReleaseSession(ctx, sessionId, opt.AcquireSessionOptions)
	}()

	// Create a new session token.
	sessionLocation := NewSessionLocation(n.Host, sessionId)
	sessionToken := NewSessionToken(sessionLocation)

	// Run the ifCreatedAndAcquired callback and return its return value.
	return sessionToken, ifCreatedAndAcquired(sessionToken)
}

// Create and acquire a session if there is no session token, otherwise acquire
// the session, then run the ifAcquired callback. Inside the callback is
// possible to safely use the session. The options defines how the session is
// created and acquired. (Note: this method is just a convenience wrapper around
// CreateAndAcquireSession and AcquireSession).
//
// There are 3 possible outcomes:
//  1. The session is acquired and the callback is run. In this case the return
//     value is the eventual error return by the callback.
//  2. The session has been offloaded and the callback is not run. In this case
//     the return value is the sessionLocation of the session.
//  3. There is an error and the callback is not run. In this case the error is
//     returned.
func (n *Node) MaybeCreateAndAcquireSession(
	ctx context.Context,
	sessionToken *SessionToken,
	opt CreateAndAcquireSessionOptions,
	ifAcquired func(SessionToken) error,
) (_ *SessionLocation, err error) {
	// If there is no session token, create and acquire a session.
	if (*sessionToken == SessionToken{}) {
		*sessionToken, err = n.CreateAndAcquireSession(ctx, opt, ifAcquired)
		// return the error if there is one.
		return nil, err
	}

	// Acquire the session.
	return n.AcquireSession(ctx, *sessionToken, opt.AcquireSessionOptions, func() error { return ifAcquired(*sessionToken) })
}
