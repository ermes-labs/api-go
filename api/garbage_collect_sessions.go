package api

import (
	"context"
)

// Commands to garbage collect sessions.
type GarbageCollectSessionsCommands interface {
	// Garbage collect sessions, the options to define how the sessions are
	// garbage collected. The function accept a cursor to continue the garbage
	// collection from the last cursor, nil to start from the beginning. The
	// function returns the next cursor to continue the garbage collection, or
	// nil if the garbage collection is completed.
	GarbageCollectSessions(
		ctx context.Context,
		opt GarbageCollectSessionsOptions,
		cursor *string,
	) (*string, error)
}

// Garbage collect sessions, the options to define how the sessions are garbage
// collected.
func (n *Node) GarbageCollectSessions(
	ctx context.Context,
	opt GarbageCollectSessionsOptions,
) error {
	var cursor *string = nil

	// Garbage collect sessions.
	for {
		// Garbage collect sessions.
		cursor, err := n.Cmd.GarbageCollectSessions(ctx, opt, cursor)

		// If there is an error, return it.
		if err != nil {
			return err
		}

		// If the garbage collection is completed, return.
		if cursor == nil {
			break
		}
	}

	// Return nil.
	return nil
}
