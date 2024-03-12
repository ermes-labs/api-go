package api

// Options that defines how a session is created.
type ResourcesUsageOptions struct{}

// Builder for ResourcesUsageOptions.
type ResourcesUsageOptionsBuilder struct {
	options ResourcesUsageOptions
}

// Create a new ResourcesUsageOptionsBuilder.
func NewResourcesUsageOptionsBuilder() *ResourcesUsageOptionsBuilder {
	return &ResourcesUsageOptionsBuilder{
		options: DefaultResourcesUsageOptions(),
	}
}

// Build the ResourcesUsageOptions.
func (builder *ResourcesUsageOptionsBuilder) Build() ResourcesUsageOptions {
	return builder.options
}

// DefaultResourcesUsageOptions returns the default options to create a session.
func DefaultResourcesUsageOptions() ResourcesUsageOptions {
	return ResourcesUsageOptions{}
}
