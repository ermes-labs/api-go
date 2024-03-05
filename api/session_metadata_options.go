package api

import (
	"time"

	"github.com/ermes-labs/api-go/infrastructure"
)

// Options that defines how a session is created.
type SessionMetadataOptions struct {
	// The geographic coordinates associated with the client that owns the session.
	// If nil, the client sessionLocation is initially approximated to the sessionLocation of
	// the node that creates the session. Default is nil.
	clientGeoCoordinates *infrastructure.GeoCoordinates
	// The expiration time is expressed as a Unix timestamp (UTC). If the
	// expiration time is nil, the session does not expire. Default is nil.
	// If the "expired" option is set this field is ignored.
	expiresAt *int64
	// If true, the session is considered expired. Default is false. If true, the
	// "expiresAt" field is ignored.
	expired bool
}

// Builder for SessionMetadataOptions.
type SessionMetadataOptionsBuilder struct {
	options SessionMetadataOptions
}

// Create a new SessionMetadataOptionsBuilder.
func NewSessionMetadataOptionsBuilder() *SessionMetadataOptionsBuilder {
	return &SessionMetadataOptionsBuilder{
		options: DefaultSessionMetadataOptions(),
	}
}

// Set the client geo coordinates.
func (builder *SessionMetadataOptionsBuilder) ClientGeoCoordinates(clientGeoCoordinates infrastructure.GeoCoordinates) *SessionMetadataOptionsBuilder {
	builder.options.clientGeoCoordinates = &clientGeoCoordinates
	return builder
}

// Set the session expiration time.
func (builder *SessionMetadataOptionsBuilder) ExpiresAt(expiresAt time.Time) *SessionMetadataOptionsBuilder {
	expiresAtUnix := expiresAt.Unix()
	builder.options.expiresAt = &expiresAtUnix
	return builder
}

// Set the session expiration time as a duration from now.
func (builder *SessionMetadataOptionsBuilder) Expires(expiresIn time.Duration) *SessionMetadataOptionsBuilder {
	expiresAtUnix := time.Now().Add(expiresIn).Unix()
	builder.options.expiresAt = &expiresAtUnix
	return builder
}

func (builder *SessionMetadataOptionsBuilder) UnixExpiresAt(expiresAt int64) *SessionMetadataOptionsBuilder {
	builder.options.expiresAt = &expiresAt
	return builder
}

// Set the session expiration time as a duration from now.
func (builder *SessionMetadataOptionsBuilder) UnixExpires(expiresIn int64) *SessionMetadataOptionsBuilder {
	expiresAtUnix := time.Now().Unix() + expiresIn
	builder.options.expiresAt = &expiresAtUnix
	return builder
}

func (builder *SessionMetadataOptionsBuilder) MarkExpired() *SessionMetadataOptionsBuilder {
	builder.options.expired = true
	return builder
}

// Build the SessionMetadataOptions.
func (builder *SessionMetadataOptionsBuilder) Build() SessionMetadataOptions {
	return builder.options
}

// DefaultSessionMetadataOptions returns the default options to create a session.
func DefaultSessionMetadataOptions() SessionMetadataOptions {
	return SessionMetadataOptions{
		clientGeoCoordinates: nil,
		expiresAt:            nil,
		expired:              false,
	}
}
