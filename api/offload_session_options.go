package api

import (
	"github.com/ermes-labs/api-go/infrastructure"
)

// Options for the offloadSession.
type OffloadSessionOptions struct {
	// The id of the session to offload.
	sessionId string
	// The sessionLocation to offload the session to.
	toLocation infrastructure.SessionLocation
}

// Builder for OffloadSessionOptions.
type OffloadSessionOptionsBuilder struct {
	options OffloadSessionOptions
}

// Create a new OffloadSessionOptionsBuilder.
func NewOffloadSessionOptionsBuilder() *OffloadSessionOptionsBuilder {
	return &OffloadSessionOptionsBuilder{
		options: DefaultOffloadSessionOptions(),
	}
}

// Set the sessionId.
func (builder *OffloadSessionOptionsBuilder) SessionId(sessionId string) *OffloadSessionOptionsBuilder {
	builder.options.sessionId = sessionId
	return builder
}

// Set the toLocation.
func (builder *OffloadSessionOptionsBuilder) ToLocation(toLocation infrastructure.SessionLocation) *OffloadSessionOptionsBuilder {
	builder.options.toLocation = toLocation
	return builder
}

// Build the OffloadSessionOptions.
func (builder *OffloadSessionOptionsBuilder) Build() OffloadSessionOptions {
	return builder.options
}

// DefaultOffloadSessionOptions returns the default options for the offloadSession.
func DefaultOffloadSessionOptions() OffloadSessionOptions {
	return OffloadSessionOptions{}
}
