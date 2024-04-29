package http_functions

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ermes-labs/api-go/api"
)

var (
	// Header names.
	lastVisitedSessionIdHeaderName = "X-Session-Last-Visited-Session-Id"
	newLocationHeaderName          = "X-Session-New-Location"
	// Type name.
	confirmOffloadRequestType = "confirm_offload"
)

func (h *Handler) ConfirmOffload(
	w http.ResponseWriter,
	req *http.Request,
) {
	// Extract the last visited session ID and the new location from the headers.
	lastVisitedSessionId := req.Header.Get(lastVisitedSessionIdHeaderName)
	newLocationString := req.Header.Get(newLocationHeaderName)
	// Unmarshal the new location.
	var newLocation api.SessionLocation
	if err := json.Unmarshal([]byte(newLocationString), &newLocation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clientRedirected, err := h.node.UpdateOffloadedSessionLocation(req.Context(), lastVisitedSessionId, newLocation)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatBool(clientRedirected)))
}

func (h *Handler) CreateConfirmOffloadRequest(
	lastVisitedLocation api.SessionLocation,
	newLocation api.SessionLocation,
) (*http.Request, error) {
	// Create the URL.
	url := url.URL{
		Scheme:   h.Scheme,
		Host:     lastVisitedLocation.Host,
		Path:     h.Path,
		RawQuery: "type=" + confirmOffloadRequestType,
	}

	// Create the request.
	req, err := http.NewRequest("POST", url.String(), nil)
	if err != nil {
		return nil, err
	}

	// Serialize the session locations.
	newLocationBytes, err := json.Marshal(newLocation)
	if err != nil {
		return nil, err
	}

	// Set the headers.
	req.Header.Set(lastVisitedSessionIdHeaderName, lastVisitedLocation.SessionId)
	req.Header.Set(newLocationHeaderName, string(newLocationBytes))
	// Return the request.
	return req, nil
}

func (h *Handler) IssueConfirmOffloadRequest(
	ctx context.Context,
	lastVisitedLocation api.SessionLocation,
	newLocation api.SessionLocation,
) (bool, error) {
	// Create the request.
	req, err := h.CreateConfirmOffloadRequest(
		lastVisitedLocation,
		newLocation,
	)

	if err != nil {
		return false, err
	}

	// Perform the request.
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return false, err
	}

	// Check the status code.
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("confirm offload request failed")
	}

	// Return the response.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	clientRedirected, err := strconv.ParseBool(string(body))
	if err != nil {
		return false, err
	}
	return clientRedirected, nil
}
