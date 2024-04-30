package http_functions

import (
	"net/http"

	"github.com/ermes-labs/api-go/api"
)

type Handler struct {
	node   *api.Node
	Scheme string
	Path   string
}

func NewHandler(node *api.Node, scheme, path string) *Handler {
	return &Handler{
		node:   node,
		Scheme: scheme,
		Path:   path,
	}
}

func (h *Handler) Handle(
	w http.ResponseWriter,
	req *http.Request,
) {
	// Defer the close of the response body.
	defer req.Body.Close()

	// Get the request type from the request query.
	requestType := req.URL.Query().Get("type")

	// Handle the request based on the request type.
	switch requestType {
	case offloadRequestType:
		h.Offload(w, req)
	case onloadRequestType:
		h.Onload(w, req)
	case confirmOffloadRequestType:
		h.ConfirmOffload(w, req)
	default:
		http.Error(w, "Invalid request type", http.StatusBadRequest)
	}
}
