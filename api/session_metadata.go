package api

import (
	"context"

	"github.com/ermes-labs/api-go/infrastructure"
)

// Metadata associated with a session.
type SessionMetadata struct {
	// The geographic coordinates associated with the client that owns the session.
	ClientGeoCoordinates *infrastructure.GeoCoordinates
	// The id of the node that created the session.
	CreatedIn string
	// The timestamp when the session was created, expressed as a Unix timestamp
	// (UTC).
	CreatedAt int64
	// The updated timestamp when the session was last updated, expressed as a
	// Unix timestamp (UTC).
	UpdatedAt int64
	// The expiration time is expressed as a Unix timestamp (UTC). If the
	// expiration time is nil, the session does not expire.
	ExpiresAt *int64
}

// Commands to manage the metadata of a session.
type SessionMetadataCommands interface {
	// Returns the metadata associated with a session.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	GetSessionMetadata(
		ctx context.Context,
		sessionId string,
	) (SessionMetadata, error)
	// SetClientCoordinates sets the coordinates of the client of a session.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	SetSessionMetadata(
		ctx context.Context,
		sessionId string,
		opt SessionMetadataOptions,
	) error
}

// Get the metadata associated with a session.
func (n *Node) GetSessionMetadata(
	ctx context.Context,
	sessionId string,
) (SessionMetadata, error) {
	return n.Cmd.GetSessionMetadata(ctx, sessionId)
}

// Set the metadata associated with a session.
func (n *Node) SetSessionMetadata(
	ctx context.Context,
	sessionId string,
	opt SessionMetadataOptions,
) error {
	return n.Cmd.SetSessionMetadata(ctx, sessionId, opt)
}
