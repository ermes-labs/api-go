package api

import (
	"net/http"
)

// Options for the handler.
type HTTPHandlerOptions struct {
	AcquireSessionOptions              func(req *http.Request) AcquireSessionOptions
	CreateSessionOptions               func(req *http.Request) CreateSessionOptions
	getSessionTokenBytes               func(req *http.Request) []byte
	setSessionTokenBytes               func(w http.ResponseWriter, sessionTokenBytes []byte)
	redirectResponse                   func(w http.ResponseWriter, req *http.Request, host string)
	malformedSessionTokenErrorResponse func(w http.ResponseWriter, err error)
	internalServerErrorResponse        func(w http.ResponseWriter, err error)
}

// Builder for HTTPHandlerOptions.
type HTTPHandlerOptionsBuilder struct {
	options HTTPHandlerOptions
}

// Create a new HTTPHandlerOptionsBuilder.
func NewHTTPHandlerOptionsBuilder() *HTTPHandlerOptionsBuilder {
	return &HTTPHandlerOptionsBuilder{
		options: DefaultHTTPHandlerOptions(),
	}
}

// Set the AcquireSessionOptions function.
func (builder *HTTPHandlerOptionsBuilder) AcquireSessionOptions(AcquireSessionOptions func(req *http.Request) AcquireSessionOptions) *HTTPHandlerOptionsBuilder {
	builder.options.AcquireSessionOptions = AcquireSessionOptions
	return builder
}

// Set the CreateSessionOptions function.
func (builder *HTTPHandlerOptionsBuilder) CreateSessionOptions(CreateSessionOptions func(req *http.Request) CreateSessionOptions) *HTTPHandlerOptionsBuilder {
	builder.options.CreateSessionOptions = CreateSessionOptions
	return builder
}

// Set the getSessionTokenBytes function.
func (builder *HTTPHandlerOptionsBuilder) GetSessionTokenBytes(getSessionTokenBytes func(req *http.Request) []byte) *HTTPHandlerOptionsBuilder {
	builder.options.getSessionTokenBytes = getSessionTokenBytes
	return builder
}

// Set the setSessionTokenBytes function.
func (builder *HTTPHandlerOptionsBuilder) SetSessionTokenBytes(setSessionTokenBytes func(w http.ResponseWriter, sessionTokenBytes []byte)) *HTTPHandlerOptionsBuilder {
	builder.options.setSessionTokenBytes = setSessionTokenBytes
	return builder
}

// Set the redirectResponse function.
func (builder *HTTPHandlerOptionsBuilder) RedirectResponse(redirectResponse func(w http.ResponseWriter, req *http.Request, host string)) *HTTPHandlerOptionsBuilder {
	builder.options.redirectResponse = redirectResponse
	return builder
}

// Set the malformedSessionTokenErrorResponse function.
func (builder *HTTPHandlerOptionsBuilder) MalformedSessionTokenErrorResponse(malformedSessionTokenErrorResponse func(w http.ResponseWriter, err error)) *HTTPHandlerOptionsBuilder {
	builder.options.malformedSessionTokenErrorResponse = malformedSessionTokenErrorResponse
	return builder
}

// Set the internalServerErrorResponse function.
func (builder *HTTPHandlerOptionsBuilder) InternalServerErrorResponse(internalServerErrorResponse func(w http.ResponseWriter, err error)) *HTTPHandlerOptionsBuilder {
	builder.options.internalServerErrorResponse = internalServerErrorResponse
	return builder
}

// Set the getSessionTokenBytes and setSessionTokenBytes functions to use the
// given header name to get and set the session token.
func (builder *HTTPHandlerOptionsBuilder) SessionTokenHeaderName(header string) *HTTPHandlerOptionsBuilder {
	builder.options.getSessionTokenBytes = func(req *http.Request) []byte {
		return GetSessionTokenBytesFromHeader(req, header)
	}
	builder.options.setSessionTokenBytes = func(w http.ResponseWriter, sessionTokenBytes []byte) {
		SetSessionTokenBytesToHeader(w, sessionTokenBytes, header)
	}
	return builder
}

// Build the HTTPHandlerOptions.
func (builder *HTTPHandlerOptionsBuilder) Build() HTTPHandlerOptions {
	return builder.options
}

// DefaultHTTPHandlerOptions returns the default options for the handler.
func DefaultHTTPHandlerOptions() HTTPHandlerOptions {
	return HTTPHandlerOptions{
		AcquireSessionOptions: func(_ *http.Request) AcquireSessionOptions {
			// Return the default options to acquire a session.
			return DefaultAcquireSessionOptions()
		},
		CreateSessionOptions: func(_ *http.Request) CreateSessionOptions {
			// Return the default options to create a session.
			return DefaultCreateSessionOptions()
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
