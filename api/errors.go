package api

import (
	"errors"
	"fmt"
)

var (
	// ErrErmes is the base error for all errors in this package.
	ErrErmes = errors.New("[ermes]")
	// ErrSessionNotFound is returned when a session is not found.
	ErrSessionNotFound = fmt.Errorf("%w: session not found", ErrErmes)
	// ErrSessionIsOffloading is returned when an action cannot be performed because a session is offloading.
	ErrSessionIsOffloading = fmt.Errorf("%w: session is offloading", ErrErmes)
	// ErrSessionAlreadyOnloaded is returned when the session is already onloaded.
	ErrSessionAlreadyOnloaded = fmt.Errorf("%w: session already onloaded", ErrErmes)
	// ErrSessionIdAlreadyExists is returned when a session id already exists.
	ErrSessionIdAlreadyExists = fmt.Errorf("%w: session id already exists", ErrErmes)
	// ErrNoAcquisitionToRelease is returned when there is no acquisition to release.
	ErrNoAcquisitionToRelease = fmt.Errorf("%w: no acquisition to release", ErrErmes)
	// ErrUnableToOffloadAcquiredSession is returned when the session is unable to offload acquired session.
	ErrUnableToOffloadAcquiredSession = fmt.Errorf("%w: unable to offload acquired session", ErrErmes)
)
