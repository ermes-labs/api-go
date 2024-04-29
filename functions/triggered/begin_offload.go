package triggered_functions

import (
	"context"

	"github.com/ermes-labs/api-go/api"
	http_functions "github.com/ermes-labs/api-go/functions/http"
)

var bestOffloadTargetsOptions = api.DefaultBestOffloadTargetsOptions()

func Begin_offload(
	node api.Node,
	ctx context.Context,
) (*api.SessionLocation, error) {
	sessions, err := node.BestSessionsToOffload(ctx, bestOffloadTargetsOptions)
	// Extract sessions ids
	sessionsIds := make([]string, 0, len(sessions))
	for sessionId := range sessions {
		sessionsIds = append(sessionsIds, sessionId)
	}

	lookupNode, err := node.FindLookupNode(ctx, sessionsIds)
	if err != nil {
		return nil, err
	}

	var sessionsToNodesMap [][2]string
	if lookupNode.Host == node.Host {
		sessionsToNodesMap, err = node.BestOffloadTargetNodes(ctx, node.Host, sessions, bestOffloadTargetsOptions)
	} else {
		var handler http_functions.Handler
		// Create the request.
		sessionsToNodesMap, err = handler.IssueBestOffloadTargetsRequest(ctx, lookupNode.Host, sessions)
	}

	if err != nil {
		return nil, err
	}

	// For each session-node couple, try to offload the session.
	for _, target := range sessionsToNodesMap {
		sessionId := target[0]
		host := target[1]

		var handler http_functions.Handler
		newSessionId, err := handler.IssueOffloadRequest(ctx, host, sessionId)

		if err != nil {
			continue
		}

		location := api.NewSessionLocation(node.Host, newSessionId)
		// musrchal in the body
		return &location, nil
	}

	return nil, nil
}
