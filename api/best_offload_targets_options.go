package api

// Options to get the best targets to offload.
type BestOffloadTargetsOptions struct {
	// The maximum number of targets to return, the system will decide which number
	// of targets is best to return, but it will not return more than this number.
	MaxTargets int
}

type BestOffloadTargetsOptionsBuilder struct {
	options BestOffloadTargetsOptions
}

// Create a new BestOffloadTargetsOptionsBuilder.
func NewBestOffloadTargetsOptionsBuilder() *BestOffloadTargetsOptionsBuilder {
	return &BestOffloadTargetsOptionsBuilder{
		options: DefaultBestOffloadTargetsOptions(),
	}
}

// Set the maximum number of targets to return.
func (builder *BestOffloadTargetsOptionsBuilder) MaxTargets(maxSessions int) *BestOffloadTargetsOptionsBuilder {
	builder.options.MaxTargets = maxSessions
	return builder
}

// Build the BestOffloadTargetsOptions.
func (builder *BestOffloadTargetsOptionsBuilder) Build() BestOffloadTargetsOptions {
	return builder.options
}

// DefaultBestOffloadTargetsOptions returns the default options to get the best
// targets to offload.
func DefaultBestOffloadTargetsOptions() BestOffloadTargetsOptions {
	return BestOffloadTargetsOptions{
		MaxTargets: 10,
	}
}
