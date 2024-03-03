package api

import "time"

// Options to start a garbage collection.
type GarbageCollectSessionsOptions struct {
	// Enable collection of expired but unreleased sessions older than the given
	// duration.
	expiredUnreleasedOlderThan *int64
}

// Builder for GarbageCollectSessionsOptions.
type GarbageCollectSessionsOptionsBuilder struct {
	options GarbageCollectSessionsOptions
}

// Create a new GarbageCollectSessionsOptionsBuilder.
func NewGarbageCollectSessionsOptionsBuilder() *GarbageCollectSessionsOptionsBuilder {
	return &GarbageCollectSessionsOptionsBuilder{
		options: DefaultGarbageCollectSessionsOptions(),
	}
}

// Enable collection of expired but unreleased sessions older than the given
// duration.
func (builder *GarbageCollectSessionsOptionsBuilder) CollectExpiredButUnreleasedOlderThan(expiredUnreleasedOlderThan time.Duration) *GarbageCollectSessionsOptionsBuilder {
	expiredUnreleasedOlderThanUnix := expiredUnreleasedOlderThan.Nanoseconds() / 1000000000
	builder.options.expiredUnreleasedOlderThan = &expiredUnreleasedOlderThanUnix
	return builder
}

// Enable collection of expired but unreleased sessions older than the given
// duration.
func (builder *GarbageCollectSessionsOptionsBuilder) CollectExpiredButUnreleasedOlderThanUnix(expiredUnreleasedOlderThan int64) *GarbageCollectSessionsOptionsBuilder {
	builder.options.expiredUnreleasedOlderThan = &expiredUnreleasedOlderThan
	return builder
}

// Build the GarbageCollectSessionsOptions.
func (builder *GarbageCollectSessionsOptionsBuilder) Build() GarbageCollectSessionsOptions {
	return builder.options
}

// DefaultGarbageCollectSessionsOptions returns the default options to acquire a
// session.
func DefaultGarbageCollectSessionsOptions() GarbageCollectSessionsOptions {
	return GarbageCollectSessionsOptions{
		expiredUnreleasedOlderThan: nil,
	}
}
