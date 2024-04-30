package http_functions

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ermes-labs/api-go/api"
)

var (
	nodeIdQueryParameterName = "nodeId"

	bestOffloadTargetsOptions = api.DefaultBestOffloadTargetsOptions()
)

func (h *Handler) BestOffloadTargets(
	w http.ResponseWriter,
	req *http.Request,
) {
	// Extract the node ID from the headers.
	nodeId := req.URL.Query().Get(nodeIdQueryParameterName)
	// Extract sessions from request body.
	var sessions map[string]api.SessionInfoForOffloadDecision
	if err := json.NewDecoder(req.Body).Decode(&sessions); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the best offload targets.
	targets, err := h.node.BestOffloadTargetNodes(req.Context(), nodeId, sessions, bestOffloadTargetsOptions)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the best offload targets.
	targetsBytes, err := json.Marshal(targets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the best offload targets.
	w.WriteHeader(http.StatusOK)
	w.Write(targetsBytes)
}

func (h *Handler) CreateBestOffloadTargetsRequest(
	nodeId string,
	sessions map[string]api.SessionInfoForOffloadDecision,
) (*http.Request, error) {
	// Create the URL.
	url := h.Scheme + "://" + h.Path

	// Create the request.
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	// Set the query parameters.
	query := req.URL.Query()
	query.Set(nodeIdQueryParameterName, nodeId)
	req.URL.RawQuery = query.Encode()

	// Serialize the sessions.
	sessionsBytes, err := json.Marshal(sessions)
	if err != nil {
		return nil, err
	}

	// Set the body.
	req.Body = io.NopCloser(bytes.NewReader(sessionsBytes))

	return req, nil
}

func (h *Handler) IssueBestOffloadTargetsRequest(
	ctx context.Context,
	nodeId string,
	sessions map[string]api.SessionInfoForOffloadDecision,
) ([][2]string, error) {
	// Create the request.
	req, err := h.CreateBestOffloadTargetsRequest(
		nodeId,
		sessions,
	)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return [][2]string{}, nil
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("best offload targets request failed")
	}

	var targets [][2]string
	if err := json.NewDecoder(res.Body).Decode(&targets); err != nil {
		return nil, err
	}

	return targets, nil
}
