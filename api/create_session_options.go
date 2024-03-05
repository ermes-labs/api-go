package api

import (
	"time"

	"github.com/ermes-labs/api-go/infrastructure"
)

// Options that defines how a session is created.
type CreateSessionOptions struct {
	// The geographic coordinates associated with the client that owns the session.
	// If nil, the client sessionLocation is initially approximated to the sessionLocation of
	// the node that creates the session. Default is nil.
	clientGeoCoordinates *infrastructure.GeoCoordinates
	// The expiration time is expressed as a Unix timestamp (UTC). If the
	// expiration time is nil, the session does not expire. Default is nil.
	expiresAt *int64
	// Optional session ID. If nil, a new session ID is generated. Default is nil.
	sessionId *string
}

// Get the geographic coordinates associated with the client that owns the session.
func (o CreateSessionOptions) ClientGeoCoordinates() *infrastructure.GeoCoordinates {
	return o.clientGeoCoordinates
}

// Get the expiration time.
func (o CreateSessionOptions) ExpiresAt() *int64 {
	return o.expiresAt
}

// Get the session ID.
func (o CreateSessionOptions) SessionId() *string {
	return o.sessionId
}

// Builder for CreateSessionOptions.
type CreateSessionOptionsBuilder struct {
	*CreateSessionOptionsBuilder
	options CreateSessionOptions
}

// Create a new CreateSessionOptionsBuilder.
func NewCreateSessionOptionsBuilder() *CreateSessionOptionsBuilder {
	return &CreateSessionOptionsBuilder{
		options: DefaultCreateSessionOptions(),
	}
}

// Set the client geo coordinates.
func (builder *CreateSessionOptionsBuilder) ClientGeoCoordinates(clientGeoCoordinates infrastructure.GeoCoordinates) *CreateSessionOptionsBuilder {
	builder.options.clientGeoCoordinates = &clientGeoCoordinates
	return builder
}

// Set the session expiration time.
func (builder *CreateSessionOptionsBuilder) ExpiresAt(expiresAt time.Time) *CreateSessionOptionsBuilder {
	expiresAtUnix := expiresAt.Unix()
	builder.options.expiresAt = &expiresAtUnix
	return builder
}

// Set the session expiration time as a duration from now.
func (builder *CreateSessionOptionsBuilder) Expires(expiresIn time.Duration) *CreateSessionOptionsBuilder {
	expiresAtUnix := time.Now().Add(expiresIn).Unix()
	builder.options.expiresAt = &expiresAtUnix
	return builder
}

func (builder *CreateSessionOptionsBuilder) UnixExpiresAt(expiresAt int64) *CreateSessionOptionsBuilder {
	builder.options.expiresAt = &expiresAt
	return builder
}

// Set the session expiration time as a duration from now.
func (builder *CreateSessionOptionsBuilder) UnixExpires(expiresIn int64) *CreateSessionOptionsBuilder {
	expiresAtUnix := time.Now().Unix() + expiresIn
	builder.options.expiresAt = &expiresAtUnix
	return builder
}

// Set the session ID.
func (builder *CreateSessionOptionsBuilder) SessionId(sessionId string) *CreateSessionOptionsBuilder {
	builder.options.sessionId = &sessionId
	return builder
}

// Build the CreateSessionOptions.
func (builder *CreateSessionOptionsBuilder) Build() CreateSessionOptions {
	return builder.options
}

// DefaultCreateSessionOptions returns the default options to create a session.
func DefaultCreateSessionOptions() CreateSessionOptions {
	return CreateSessionOptions{
		clientGeoCoordinates: nil,
		expiresAt:            nil,
		sessionId:            nil,
	}
}
