package api

import (
	"context"
)

// Commands to acquire and release sessions.
type AcquireSessionCommands interface {
	// Acquires a session. If the session has been offloaded and not acquired it
	// returns the new session sessionLocation, otherwise nil. The options defines how
	// the session is acquired.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	// - ErrSessionIsOffloading: If the session is offloading and the required permission is read-write.
	AcquireSession(
		ctx context.Context,
		sessionId string,
		opt AcquireSessionOptions,
	) (*SessionLocation, error)
	// Releases a previously acquired session. The options defines how the session
	// is released.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	// - ErrNoAcquisitionToRelease: If there is no acquisition to release.
	ReleaseSession(
		ctx context.Context,
		sessionId string,
		opt AcquireSessionOptions,
	) (*SessionLocation, error)
}

// Acquire a session, then run the ifAcquired callback. Inside the callback is
// possible to safely use the session key space with redis. The options defines
// how the session is acquired.
//
// There are 3 possible outcomes:
//  1. The session is acquired and the callback is run. In this case the return
//     value is nil.
//  2. The session has been offloaded and the callback is not run. In this case
//     the return value is the sessionLocation of the session.
//  3. There is an error and the callback is not run. In this case the error is
//     returned (e.g. when the session is offloading).
func (n *Node) AcquireSession(
	cmd AcquireSessionCommands,
	ctx context.Context,
	sessionToken SessionToken,
	opt AcquireSessionOptions,
	ifAcquired func() error,
) (*SessionLocation, error) {
	offloadedTo, err := cmd.AcquireSession(ctx, sessionToken.SessionId, opt)

	// If there is an error, return it.
	if err != nil {
		return nil, err
	}

	// If the session has been offloaded, return the sessionLocation of the session.
	if offloadedTo != nil {
		return offloadedTo, nil
	}

	// Defer the release of the session metadata.
	defer func() {
		cmd.ReleaseSession(ctx, sessionToken.SessionId, opt)
	}()

	// Run the ifAcquired callback and return its return value.
	return nil, ifAcquired()
}
