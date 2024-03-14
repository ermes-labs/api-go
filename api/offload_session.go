package api

import (
	"context"
	"io"
)

type OffloadSessionCommands interface {
	SessionMetadataCommands
	// OffloadStart starts the offload of a session. The function returns the
	// io.Reader that allows to read the session data, an optional loader function
	// to fulfill the io.Reader, and an error. The function is thought to be
	// used in scenarios where the session data is huge and streaming is
	// required. The loader function will be run concurrently to the reader process.
	// Errors can flow from the loader function to the reader passing trough the
	// io.Reader, vice-versa the loader should stop if the context is canceled.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	// - ErrSessionIsOffloading: If the session is already offloading.
	// - ErrUnableToOffloadAcquiredSession: If the session is unable to offload because it is acquired.
	OffloadSession(
		ctx context.Context,
		id string,
		opt OffloadSessionOptions,
	) (sessionDataReadCloser io.ReadCloser, loader func(), err error)
	// Confirms the offload of a session.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	ConfirmSessionOffload(
		ctx context.Context,
		id string,
		newLocation SessionLocation,
		opt OffloadSessionOptions,
		// TODO: extract into another API.
		notifyLastVisitedNode func(ctx context.Context, oldLocation SessionLocation) (clientRedirected bool, err error),
	) (err error)
	// Updates the location of an offloaded session, the function returns true if
	// the client has already been redirected to the new location, while the update
	// is in progress. If true, this node is no more the last visited one, otherwise
	// the node is still the last visited one.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	// - ErrSessionIsNotOffloaded: If the session is not offloaded.
	UpdateOffloadedSessionLocation(
		ctx context.Context,
		id string,
		newLocation SessionLocation,
	) (clientRedirected bool, err error)
	// Returns the offloaded sessions, the function returns the new cursor, the
	// list of session ids and an error. The cursor is used to paginate the results.
	// If the cursor is empty, the function returns the first page of results.
	// errors:
	// - ErrInvalidCursor: If the cursor is invalid.
	// - ErrInvalidCount: If the count is invalid.
	ScanOffloadedSessions(
		ctx context.Context,
		cursor uint64,
		count int64,
	) (ids []string, newCursor uint64, err error)
}

// Offloads a session to a new location. The function returns the new location of
// the session.
func (n *Node) OffloadSession(
	ctx context.Context,
	sessionId string,
	opt OffloadSessionOptions,
	onload func(ctx context.Context, metadata SessionMetadata, reader io.Reader) (SessionLocation, error),
	notifyLastVisitedNode func(ctx context.Context, oldLocation SessionLocation, newLocation SessionLocation) (bool, error),
) (SessionLocation, error) {
	// Create a new context to cancel the loader if the context is canceled.
	ctx, cancel := context.WithCancel(ctx)

	// Read the metadata of the session.
	metadata, err := n.cmd.GetSessionMetadata(ctx, sessionId)
	// If there is an error, return it.
	if err != nil {
		cancel()
		return SessionLocation{}, err
	}

	// Start the offload of the session.
	reader, loader, err := n.cmd.OffloadSession(ctx, sessionId, opt)
	// If there is an error, return it.
	if err != nil {
		cancel()
		return SessionLocation{}, err
	}

	// If the loader failed, default to false.
	loaderFailed := false
	// If there is a loader, run it concurrently in a go routine.
	if loader != nil {
		// Wrap the reader
		reader = wrapReaderWithCheck(reader, func(err error) {
			if err != io.EOF {
				// If the loader failed, cancel the context.
				loaderFailed = true
				cancel()
			}
		})

		// Run the loader concurrently.
		go loader()
	}

	// Run the onload function.
	newLocation, err := onload(ctx, metadata, reader)
	// We could close them only in case of error, but we do it always to be sure.
	cancel()
	reader.Close()
	// If there is an error, return it.
	if err != nil {
		return SessionLocation{}, err
	}

	// If there was an error during the streaming process but for some reason the
	// onloading node confirmed the offload.
	if loaderFailed {
		// TODO: What to do here?
	} else {
		// TODO: DO we assume the loader finished?
	}

	// Confirm the offload of the session.
	err = n.cmd.ConfirmSessionOffload(ctx, sessionId, newLocation, opt, func(ctx context.Context, oldLocation SessionLocation) (bool, error) {
		return notifyLastVisitedNode(ctx, oldLocation, newLocation)
	})
	// If there is an error, return it.
	if err != nil {
		// TODO: What to do here?
	}

	// Return the metadata and the reader.
	return newLocation, nil
}

func (n *Node) UpdateOffloadedSessionLocation(
	ctx context.Context,
	id string,
	newLocation SessionLocation,
) (clientRedirected bool, err error) {
	return n.cmd.UpdateOffloadedSessionLocation(ctx, id, newLocation)
}

func (n *Node) ScanOffloadedSessions(
	ctx context.Context,
	cursor uint64,
	count int64,
) (ids []string, newCursor uint64, err error) {
	return n.cmd.ScanOffloadedSessions(ctx, cursor, count)
}

// Wrap the readCloser with a check function.
type errorCheckingReadCloser struct {
	io.ReadCloser
	readCloser io.ReadCloser
	onError    func(error)
}

// Read from the readCloser and check for errors.
func (r *errorCheckingReadCloser) Read(p []byte) (n int, err error) {
	n, err = r.readCloser.Read(p)
	if err != nil && err != io.EOF {
		r.onError(err)
	}

	return n, err
}

// Close the readCloser.
func (r *errorCheckingReadCloser) Close() error {
	return r.readCloser.Close()
}

// Wrap the reader with a check function.
func wrapReaderWithCheck(readCloser io.ReadCloser, onError func(error)) io.ReadCloser {
	return &errorCheckingReadCloser{readCloser: readCloser, onError: onError}
}
