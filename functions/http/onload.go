package http_functions

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/ermes-labs/api-go/api"
)

var (
	// Header names.
	oldLocationHeaderName = "X-Session-Old-Location"
	metadataHeaderName    = "X-Session-Metadata"
	// Options.
	onloadOptions = api.NewOnloadSessionOptionsBuilder().Build()
	// Type name.
	onloadRequestType = "onload"
)

func (h *Handler) Onload(
	w http.ResponseWriter,
	req *http.Request,
) {
	// FIXME: should this go inside sessionMetadata?
	var oldLocation api.SessionLocation
	// Read the old location from the headers.
	oldLocationString := req.Header.Get(oldLocationHeaderName)
	// Unmarshall the old location.
	if err := json.Unmarshal([]byte(oldLocationString), &oldLocation); err != nil {
		// Return an error response
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var metadata api.SessionMetadata
	// Read the metadata from the headers.
	metadataString := req.Header.Get(metadataHeaderName)
	// Unmarshall the metadata.
	err := json.Unmarshal([]byte(metadataString), &metadata)
	// If there is an error, return it.
	if err != nil {
		// Return an error response
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return the metadata and the response body.
	onloadedTo, err := h.node.OnloadSession(
		req.Context(),
		metadata,
		req.Body,
		onloadOptions,
	)

	if err == nil {
		var locationBytes []byte
		// Marshall the location.
		if locationBytes, err = json.Marshal(onloadedTo); err == nil {
			// Set the status code in the response.
			w.WriteHeader(http.StatusCreated)
			// Write the location in the response body.
			w.Write(locationBytes)
		}

		// FIXME: Handle the case in which the Marshall fails but the onload is
		// successful.
	}

	// If there is an error, return it.
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (h *Handler) CreateOnloadRequest(
	ctx context.Context,
	onloadToHost string,
	oldLocation api.SessionLocation,
	metadata api.SessionMetadata,
	body io.Reader,
) (*http.Request, error) {
	// The endpoint.
	url := url.URL{
		Scheme:   h.Scheme,
		Host:     onloadToHost,
		Path:     h.Path,
		RawQuery: "type=" + onloadRequestType,
	}

	// Create the request.
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}

	// Marshall the old location and the metadata.
	oldLocationJSON, err := json.Marshal(oldLocation)
	if err != nil {
		return nil, err
	}
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	// Set the content type in the headers.
	req.Header.Set("Content-Type", "application/octet-stream")
	// Set the old location in the headers.
	req.Header.Set(oldLocationHeaderName, string(oldLocationJSON))
	// Set the metadata in the headers.
	req.Header.Set(metadataHeaderName, string(metadataJSON))

	// Return the request.
	return req, nil
}

func (h *Handler) IssueOnloadRequest(
	ctx context.Context,
	onloadToHost string,
	oldLocation api.SessionLocation,
	metadata api.SessionMetadata,
	body io.Reader,
) (api.SessionLocation, error) {
	// Create the request.
	req, err := h.CreateOnloadRequest(ctx, onloadToHost, oldLocation, metadata, body)
	if err != nil {
		return api.SessionLocation{}, err
	}

	// Send the request.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return api.SessionLocation{}, err
	}
	// Defer the close of the response body.
	defer res.Body.Close()

	// If the status code is not OK, return an error.
	if res.StatusCode != http.StatusOK {
		// TODO: Return a more meaningful error.
		return api.SessionLocation{}, errors.New("onload failed")
	}

	// Unmarshall the location.
	var location api.SessionLocation
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		// FIXME: Handle the case in which the unmarshall fails but the onload is
		// successful.
		return api.SessionLocation{}, err
	}

	// Return the location.
	return location, nil
}
