package http

import (
	"net/http"

	"github.com/ermes-labs/api-go/api"
)

// Options for the handler.
type HandlerOptions struct {
	AcquireSessionOptions              func(req *http.Request) api.AcquireSessionOptions
	CreateSessionOptions               func(req *http.Request) api.CreateSessionOptions
	getSessionTokenBytes               func(req *http.Request) []byte
	setSessionTokenBytes               func(w http.ResponseWriter, sessionTokenBytes []byte)
	redirectResponse                   func(w http.ResponseWriter, req *http.Request, host string)
	malformedSessionTokenErrorResponse func(w http.ResponseWriter, err error)
	internalServerErrorResponse        func(w http.ResponseWriter, err error)
}

// Builder for HandlerOptions.
type HandlerOptionsBuilder struct {
	options HandlerOptions
}

// Create a new HandlerOptionsBuilder.
func NewHandlerOptionsBuilder() *HandlerOptionsBuilder {
	return &HandlerOptionsBuilder{
		options: DefaultHandlerOptions(),
	}
}

// Set the AcquireSessionOptions function.
func (builder *HandlerOptionsBuilder) AcquireSessionOptions(AcquireSessionOptions func(req *http.Request) api.AcquireSessionOptions) *HandlerOptionsBuilder {
	builder.options.AcquireSessionOptions = AcquireSessionOptions
	return builder
}

// Set the CreateSessionOptions function.
func (builder *HandlerOptionsBuilder) CreateSessionOptions(CreateSessionOptions func(req *http.Request) api.CreateSessionOptions) *HandlerOptionsBuilder {
	builder.options.CreateSessionOptions = CreateSessionOptions
	return builder
}

// Set the getSessionTokenBytes function.
func (builder *HandlerOptionsBuilder) GetSessionTokenBytes(getSessionTokenBytes func(req *http.Request) []byte) *HandlerOptionsBuilder {
	builder.options.getSessionTokenBytes = getSessionTokenBytes
	return builder
}

// Set the setSessionTokenBytes function.
func (builder *HandlerOptionsBuilder) SetSessionTokenBytes(setSessionTokenBytes func(w http.ResponseWriter, sessionTokenBytes []byte)) *HandlerOptionsBuilder {
	builder.options.setSessionTokenBytes = setSessionTokenBytes
	return builder
}

// Set the redirectResponse function.
func (builder *HandlerOptionsBuilder) RedirectResponse(redirectResponse func(w http.ResponseWriter, req *http.Request, host string)) *HandlerOptionsBuilder {
	builder.options.redirectResponse = redirectResponse
	return builder
}

// Set the malformedSessionTokenErrorResponse function.
func (builder *HandlerOptionsBuilder) MalformedSessionTokenErrorResponse(malformedSessionTokenErrorResponse func(w http.ResponseWriter, err error)) *HandlerOptionsBuilder {
	builder.options.malformedSessionTokenErrorResponse = malformedSessionTokenErrorResponse
	return builder
}

// Set the internalServerErrorResponse function.
func (builder *HandlerOptionsBuilder) InternalServerErrorResponse(internalServerErrorResponse func(w http.ResponseWriter, err error)) *HandlerOptionsBuilder {
	builder.options.internalServerErrorResponse = internalServerErrorResponse
	return builder
}

// Set the getSessionTokenBytes and setSessionTokenBytes functions to use the
// given header name to get and set the session token.
func (builder *HandlerOptionsBuilder) SessionTokenHeaderName(header string) *HandlerOptionsBuilder {
	builder.options.getSessionTokenBytes = func(req *http.Request) []byte {
		return GetSessionTokenBytesFromHeader(req, header)
	}
	builder.options.setSessionTokenBytes = func(w http.ResponseWriter, sessionTokenBytes []byte) {
		SetSessionTokenBytesToHeader(w, sessionTokenBytes, header)
	}
	return builder
}

// Build the HandlerOptions.
func (builder *HandlerOptionsBuilder) Build() HandlerOptions {
	return builder.options
}

// DefaultHandlerOptions returns the default options for the handler.
func DefaultHandlerOptions() HandlerOptions {
	return HandlerOptions{
		AcquireSessionOptions: func(_ *http.Request) api.AcquireSessionOptions {
			// Return the default options to acquire a session.
			return api.DefaultAcquireSessionOptions()
		},
		CreateSessionOptions: func(_ *http.Request) api.CreateSessionOptions {
			// Return the default options to create a session.
			return api.DefaultCreateSessionOptions()
		},
		getSessionTokenBytes: func(req *http.Request) []byte {
			return GetSessionTokenBytesFromHeader(req, DefaultTokenHeaderName)
		},
		setSessionTokenBytes: func(w http.ResponseWriter, sessionTokenBytes []byte) {
			SetSessionTokenBytesToHeader(w, sessionTokenBytes, DefaultTokenHeaderName)
		},
		redirectResponse: func(w http.ResponseWriter, req *http.Request, host string) {
			// Redirect the request to the given host.
			http.Redirect(w, req, host, http.StatusFound)
		},
		malformedSessionTokenErrorResponse: func(w http.ResponseWriter, err error) {
			// Return a bad request response with the error message.
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		internalServerErrorResponse: func(w http.ResponseWriter, err error) {
			// Return an internal server error response with the error message.
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}
}

// DefaultTokenHeaderName is the default name of the header that contains the
// session token.
const DefaultTokenHeaderName = "X-Ermes-Token"

// GetSessionTokenBytesFromHeader returns the session token bytes from the
// header of the request.
func GetSessionTokenBytesFromHeader(req *http.Request, headerName string) []byte {
	return []byte(req.Header.Get(headerName))
}

// SetSessionTokenBytesToHeader sets the session token bytes to the header of
// the response.
func SetSessionTokenBytesToHeader(w http.ResponseWriter, sessionTokenBytes []byte, headerName string) {
	w.Header().Set(headerName, string(sessionTokenBytes))
}
