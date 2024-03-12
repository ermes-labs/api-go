package api

// Options to get the best sessions to offload.
type BestSessionsToOffloadOptions struct {
	// The maximum number of sessions to return, the system will decide which number
	// of sessions is best to return, but it will not return more than this number.
	MaxSessions int
}

type BestSessionsToOffloadOptionsBuilder struct {
	options BestSessionsToOffloadOptions
}

// Create a new BestSessionsToOffloadOptionsBuilder.
func NewBestSessionsToOffloadOptionsBuilder() *BestSessionsToOffloadOptionsBuilder {
	return &BestSessionsToOffloadOptionsBuilder{
		options: DefaultBestSessionsToOffloadOptions(),
	}
}

// Set the maximum number of sessions to return.
func (builder *BestSessionsToOffloadOptionsBuilder) MaxSessions(maxSessions int) *BestSessionsToOffloadOptionsBuilder {
	builder.options.MaxSessions = maxSessions
	return builder
}

// Build the BestSessionsToOffloadOptions.
func (builder *BestSessionsToOffloadOptionsBuilder) Build() BestSessionsToOffloadOptions {
	return builder.options
}

// DefaultBestSessionsToOffloadOptions returns the default options to get the best
// sessions to offload.
func DefaultBestSessionsToOffloadOptions() BestSessionsToOffloadOptions {
	return BestSessionsToOffloadOptions{
		MaxSessions: 10,
	}
}
