package api

// Options to get the best sessions to offload.
type bestSessionsToOffloadOptions struct {
	// The maximum number of sessions to return, the system will decide which number
	// of sessions is best to return, but it will not return more than this number.
	MaxSessions int
}

type bestSessionsToOffloadOptionsBuilder struct {
	options bestSessionsToOffloadOptions
}

// Create a new bestSessionsToOffloadOptionsBuilder.
func NewBestSessionsToOffloadOptionsBuilder() *bestSessionsToOffloadOptionsBuilder {
	return &bestSessionsToOffloadOptionsBuilder{
		options: DefaultBestSessionsToOffloadOptions(),
	}
}

// Set the maximum number of sessions to return.
func (builder *bestSessionsToOffloadOptionsBuilder) MaxSessions(maxSessions int) *bestSessionsToOffloadOptionsBuilder {
	builder.options.MaxSessions = maxSessions
	return builder
}

// Build the bestSessionsToOffloadOptions.
func (builder *bestSessionsToOffloadOptionsBuilder) Build() bestSessionsToOffloadOptions {
	return builder.options
}

// DefaultBestSessionsToOffloadOptions returns the default options to get the best
// sessions to offload.
func DefaultBestSessionsToOffloadOptions() bestSessionsToOffloadOptions {
	return bestSessionsToOffloadOptions{
		MaxSessions: 10,
	}
}
