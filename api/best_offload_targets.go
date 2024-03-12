package api

import "context"

// The information related to a session that are used to decide the best offload
// targets.
type SessionInfoForOffloadDecision struct {
	Metadata       SessionMetadata `json:"metadata"`
	ResourcesUsage ResourcesUsage  `json:"resourcesUsage"`
}

type BestOffloadTargetsCommands interface {
	// Return the best offload targets composed by the session id and the node id.
	// The options defines how the sessions are selected. Note that sessions and
	// nodes may appear multiple times in the result, to allow for multiple choices
	// of offload targets. those are not grouped by session id or node id to allow
	// to express the priority of the offload targets.
	BestOffloadTargetNodes(
		ctx context.Context,
		sessions map[string]SessionMetadata,
		opt BestSessionsToOffloadOptions,
	) ([][2]string, error)
	// Return the best sessions to offload. This list is composed by the session
	// chosen given the local context of the node (direct or indirect knowledge of
	// the status of the system).
	BestSessionsToOffload(
		ctx context.Context,
	) (sessions map[string]SessionInfoForOffloadDecision, err error)
}

// Return the best offload targets composed by the session id and the node id.
// The options defines how the sessions are selected. Note that sessions and
// nodes may appear multiple times in the result, to allow for multiple choices
// of offload targets. those are not grouped by session id or node id to allow
// to express the priority of the offload targets.
func (n *Node) BestOffloadTargetNodes(
	ctx context.Context,
	sessions map[string]SessionMetadata,
	opt BestSessionsToOffloadOptions,
) ([][2]string, error) {
	return n.cmd.BestOffloadTargetNodes(ctx, sessions, opt)
}

// Return the best sessions to offload. This list is composed by the session
// chosen given the local context of the node (direct or indirect knowledge of
// the status of the system).
func (n *Node) BestSessionsToOffload(
	ctx context.Context,
) (sessions map[string]SessionInfoForOffloadDecision, err error) {
	return n.cmd.BestSessionsToOffload(ctx)
}
