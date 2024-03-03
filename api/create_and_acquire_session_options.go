package api

// Options that defines how a session is created and acquired.
type CreateAndAcquireSessionOptions struct {
	// The options to create the session.
	CreateSessionOptions
	// The options to acquire the session.
	AcquireSessionOptions
}

// Builder for CreateAndAcquireSessionOptions.
type CreateAndAcquireSessionOptionsBuilder struct {
	*AcquireSessionOptionsBuilder
	*CreateSessionOptionsBuilder
}

// Create a new CreateAndAcquireSessionOptionsBuilder.
func NewCreateAndAcquireSessionOptionsBuilder() *CreateAndAcquireSessionOptionsBuilder {
	return &CreateAndAcquireSessionOptionsBuilder{
		AcquireSessionOptionsBuilder: NewAcquireSessionOptionsBuilder(),
		CreateSessionOptionsBuilder:  NewCreateSessionOptionsBuilder(),
	}
}

// Build the CreateAndAcquireSessionOptions.
func (builder *CreateAndAcquireSessionOptionsBuilder) Build() CreateAndAcquireSessionOptions {
	return CreateAndAcquireSessionOptions{
		CreateSessionOptions:  builder.CreateSessionOptionsBuilder.Build(),
		AcquireSessionOptions: builder.AcquireSessionOptionsBuilder.Build(),
	}
}

func DefaultCreateAndAcquireSessionOptions() CreateAndAcquireSessionOptions {
	return CreateAndAcquireSessionOptions{
		CreateSessionOptions:  DefaultCreateSessionOptions(),
		AcquireSessionOptions: DefaultAcquireSessionOptions(),
	}
}
