package api

import (
	"errors"
	"fmt"
)

var (
	// ErrErmes is the base error for all errors in this package.
	ErrErmes = errors.New("[Ermes]")
	// ErrSessionNotFound is returned when a session is not found.
	ErrSessionNotFound = fmt.Errorf("%w: session not found", ErrErmes)
	// ErrExpirationTimeNotValid is returned when the expiration time is not valid.
	ErrSessionIsOffloading = fmt.Errorf("%w: session is offloading", ErrErmes)
	// ErrNoAcquisitionToRelease is returned when there is no acquisition to release.
	ErrNoAcquisitionToRelease = fmt.Errorf("%w: no acquisition to release", ErrErmes)
	// ErrSessionAlreadyOnloaded is returned when the session is already onloaded.
	ErrSessionAlreadyOnloaded = fmt.Errorf("%w: session already onloaded", ErrErmes)
)
