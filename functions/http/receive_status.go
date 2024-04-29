package http_functions

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type payload struct {
	// Node.
	Sessions               uint                          `json:"sessions"`
	ResourcesUsageNodesMap map[string]map[string]float64 `json:"resourcesUsageNodesMap"`
}

func (h *Handler) ReceiveStatus(
	w http.ResponseWriter,
	req *http.Request,
) {
	// Extract payload from body.
	var payload payload
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update resources usage.
	err := h.node.ResourcesUsageUpdateFromChild(req.Context(), payload.Sessions, payload.ResourcesUsageNodesMap)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return ok.
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateReceiveStatusRequest(
	sessions uint,
	resourcesUsageNodesMap map[string]map[string]float64,
) (*http.Request, error) {
	// Create the URL.
	url := h.Scheme + "://" + h.Path

	// Create the request.
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	// Serialize the payload.
	payloadBytes, err := json.Marshal(payload{
		Sessions:               sessions,
		ResourcesUsageNodesMap: resourcesUsageNodesMap,
	})
	if err != nil {
		return nil, err
	}

	// Set the body.
	req.Body = io.NopCloser(bytes.NewReader(payloadBytes))
	// Return the request.
	return req, nil
}

func (h *Handler) IssueReceiveStatusRequest(
	ctx context.Context,
	sessions uint,
	resourcesUsageNodesMap map[string]map[string]float64,
) error {
	// Create the request.
	req, err := h.CreateReceiveStatusRequest(
		sessions,
		resourcesUsageNodesMap,
	)

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("receive status request failed")
	}

	return nil
}
