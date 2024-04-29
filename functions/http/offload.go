package http_functions

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/ermes-labs/api-go/api"
)

var (
	// Param names.
	offloadToHostQueryParameterName      = "toHost"
	sessionToOffloadIdQueryParameterName = "sessionId"
	// Options.
	offloadOptions = api.NewOffloadSessionOptionsBuilder().Build()
	// Type name.
	offloadRequestType = "offload"
)

func (h *Handler) Offload(
	w http.ResponseWriter,
	req *http.Request,
) {
	// Extract sessionId and NewLocation from headers
	offloadToHost := req.URL.Query().Get(offloadToHostQueryParameterName)
	oldLocation := api.SessionLocation{
		Host:      h.node.Host,
		SessionId: req.URL.Query().Get(sessionToOffloadIdQueryParameterName),
	}

	// Offload the session
	newLocation, err := h.node.OffloadSession(
		req.Context(),
		oldLocation.SessionId,
		offloadOptions,
		func(ctx context.Context, sm api.SessionMetadata, r io.Reader) (api.SessionLocation, error) {
			return h.IssueOnloadRequest(ctx, offloadToHost, oldLocation, sm, r)
		},
		func(ctx context.Context, oldLocation api.SessionLocation, newLocation api.SessionLocation) (bool, error) {
			return h.IssueConfirmOffloadRequest(ctx, oldLocation, newLocation)
		},
	)

	// If there is an error, return it.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return ok.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(newLocation.SessionId))
}

func (h *Handler) CreateOffloadRequest(
	host string,
	sessionId string,
) (*http.Request, error) {
	queryParams := url.Values{
		offloadToHostQueryParameterName:      {host},
		sessionToOffloadIdQueryParameterName: {sessionId},
		"type":                               {offloadRequestType},
	}

	url := url.URL{
		Scheme:   h.Scheme,
		Host:     host,
		Path:     h.Path,
		RawQuery: queryParams.Encode(),
	}

	return http.NewRequest(
		"POST",
		url.String(),
		nil,
	)
}

func (h *Handler) IssueOffloadRequest(
	ctx context.Context,
	host string,
	sessionId string,
) (string, error) {
	req, err := h.CreateOffloadRequest(host, sessionId)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("offload request failed")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
