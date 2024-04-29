package api

import (
	"context"

	"github.com/ermes-labs/api-go/infrastructure"
)

// The information related to a session that are used to decide the best offload
// targets.
type SessionInfoForOffloadDecision struct {
	Metadata       SessionMetadata `json:"metadata"`
	ResourcesUsage ResourcesUsage  `json:"resourcesUsage"`
}

type BestOffloadTargetsCommands interface {
	// Return the best sessions to offload. This list is composed by the session
	// chosen given the local context of the node (direct or indirect knowledge of
	// the status of the system).
	BestSessionsToOffload(
		ctx context.Context,
		opt BestOffloadTargetsOptions,
	) (sessions map[string]SessionInfoForOffloadDecision, err error)
	// Return the best offload targets composed by the session id and the node id.
	// The options defines how the sessions are selected. Note that sessions and
	// nodes may appear multiple times in the result, to allow for multiple choices
	// of offload targets. those are not grouped by session id or node id to allow
	// to express the priority of the offload targets.
	BestOffloadTargetNodes(
		ctx context.Context,
		nodeId string,
		sessions map[string]SessionInfoForOffloadDecision,
		opt BestOffloadTargetsOptions,
	) ([][2]string, error)
	// Get the lookup node for a session offloading.
	FindLookupNode(
		ctx context.Context,
		sessionIds []string,
	) (node infrastructure.Node, err error)
}

// Return the best offload targets composed by the session id and the node id.
// The options defines how the sessions are selected. Note that sessions and
// nodes may appear multiple times in the result, to allow for multiple choices
// of offload targets. those are not grouped by session id or node id to allow
// to express the priority of the offload targets.
func (n *Node) BestOffloadTargetNodes(
	ctx context.Context,
	nodeId string,
	sessions map[string]SessionInfoForOffloadDecision,
	opt BestOffloadTargetsOptions,
) ([][2]string, error) {
	return n.Cmd.BestOffloadTargetNodes(ctx, nodeId, sessions, opt)
}

// Return the best sessions to offload. This list is composed by the session
// chosen given the local context of the node (direct or indirect knowledge of
// the status of the system).
func (n *Node) BestSessionsToOffload(
	ctx context.Context,
	opt BestOffloadTargetsOptions,
) (sessions map[string]SessionInfoForOffloadDecision, err error) {
	return n.Cmd.BestSessionsToOffload(ctx, opt)
}

// Get the lookup node for a session offloading.
func (n *Node) FindLookupNode(
	ctx context.Context,
	sessionIds []string,
) (node infrastructure.Node, err error) {
	return n.Cmd.FindLookupNode(ctx, sessionIds)
}
