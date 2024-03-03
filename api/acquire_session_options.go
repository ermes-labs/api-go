package api

// Options to acquire a session.
type AcquireSessionOptions struct {
	// If true, the session will be eligible for offloading even if the session is
	// acquired.
	allowOffloading bool
	// If true, the session will be acquired even if the session is offloading
	allowWhileOffloading bool
}

// Get the value of allowOffloading.
func (o AcquireSessionOptions) AllowOffloading() bool {
	return o.allowOffloading
}

// Get the value of allowWhileOffloading.
func (o AcquireSessionOptions) AllowWhileOffloading() bool {
	return o.allowWhileOffloading
}

// Builder for CreateSessionOptions.
type AcquireSessionOptionsBuilder struct {
	options AcquireSessionOptions
}

// Create a new AcquireSessionOptionsBuilder.
func NewAcquireSessionOptionsBuilder() *AcquireSessionOptionsBuilder {
	return &AcquireSessionOptionsBuilder{
		options: DefaultAcquireSessionOptions(),
	}
}

func (builder *AcquireSessionOptionsBuilder) AllowOffloading() *AcquireSessionOptionsBuilder {
	builder.options.allowOffloading = true
	return builder
}

// Will acquire the session with read-only permissions.
func (builder *AcquireSessionOptionsBuilder) AllowWhileOffloading() *AcquireSessionOptionsBuilder {
	builder.options.allowWhileOffloading = true
	return builder
}

// Build the AcquireSessionOptions.
func (builder *AcquireSessionOptionsBuilder) Build() AcquireSessionOptions {
	return builder.options
}

// DefaultAcquireSessionOptions returns the default options to acquire a
// session.
func DefaultAcquireSessionOptions() AcquireSessionOptions {
	return AcquireSessionOptions{
		allowWhileOffloading: false,
	}
}
