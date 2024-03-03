package api

// Options to onload a session.
type OnloadSessionOptions struct{}

// Builder for CreateSessionOptions.
type OnloadSessionOptionsBuilder struct {
	options OnloadSessionOptions
}

// Create a new OnloadSessionOptionsBuilder.
func NewOnloadSessionOptionsBuilder() *OnloadSessionOptionsBuilder {
	return &OnloadSessionOptionsBuilder{
		options: DefaultOnloadSessionOptions(),
	}
}

// Build the OnloadSessionOptions.
func (builder *OnloadSessionOptionsBuilder) Build() OnloadSessionOptions {
	return builder.options
}

// DefaultOnloadSessionOptions returns the default options to onload a
// session.
func DefaultOnloadSessionOptions() OnloadSessionOptions {
	return OnloadSessionOptions{}
}
