package api

import (
	"context"
)

type ResourcesUsage = map[string]float64
type ResourcesUsageIndex = map[string]float64

// Commands to get and update the resources usage of the sessions and the nodes.
type ResourcesUsageCommands interface {
	// Get the resources usage of a session.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	GetSessionResourcesUsage(
		ctx context.Context,
		sessionId string,
	) (resourcesUsage ResourcesUsage, err error)
	// Get the resources usage of a node.
	GetNodeResourcesUsage(
		ctx context.Context,
		nodeId string,
	) (sessions uint, resourcesUsage ResourcesUsage, err error)
	// Update the resources usage of a session, this will also update the resources
	// usage of the node.
	// errors:
	// - ErrSessionNotFound: If no session with the given id is found.
	UpdateSessionResourcesUsage(
		ctx context.Context,
		sessionId string,
		resourcesUsage ResourcesUsage,
	) (err error)
	// Get the update to send to the parent node.
	ResourcesUsageUpdateToParent(
		ctx context.Context,
	) (host string, sessions uint, resourcesUsageNodesMap map[string]ResourcesUsage, err error)
	// Get the update from the child nodes.
	ResourcesUsageUpdateFromChild(
		ctx context.Context,
		sessions uint,
		resourcesUsageNodesMap map[string]ResourcesUsage,
	) (err error)
	// Redirect new requests to the best offload target.
	RedirectNewRequests(
		ctx context.Context,
	) (redirect bool, host string)
}

// Get the resources usage of a session.
// errors:
// - ErrSessionNotFound: If no session with the given id is found.
func (n *Node) GetSessionResourcesUsage(
	ctx context.Context,
	sessionId string,
) (resourcesUsage ResourcesUsage, err error) {
	return n.cmd.GetSessionResourcesUsage(ctx, sessionId)
}

// Get the resources usage of a node.
func (n *Node) GetNodeResourcesUsage(
	ctx context.Context,
	nodeId string,
) (sessions uint, resourcesUsage ResourcesUsage, err error) {
	return n.cmd.GetNodeResourcesUsage(ctx, nodeId)
}

// Get the resources usage of all the nodes.
func (n *Node) GetNodeResourcesUsageIndex(
	ctx context.Context,
	nodeId string,
) (resourcesUsageIndex ResourcesUsageIndex, err error) {
	_, resourcesUsage, err := n.cmd.GetNodeResourcesUsage(ctx, nodeId)

	if err != nil {
		return nil, err
	}

	resourcesUsageIndex = make(ResourcesUsageIndex)
	// for each resource in the node, compute the index (value/usage)
	for resource, value := range n.Resources {
		resourcesUsageIndex[resource] = value / resourcesUsage[resource]
	}

	return resourcesUsageIndex, nil
}

// Update the resources usage of a session, this will also update the resources
// usage of the node.
// errors:
// - ErrSessionNotFound: If no session with the given id is found.
func (n *Node) UpdateSessionResourcesUsage(
	ctx context.Context,
	sessionId string,
	resourcesUsage ResourcesUsage,
) (err error) {
	return n.cmd.UpdateSessionResourcesUsage(ctx, sessionId, resourcesUsage)
}

// Get the update to send to the parent node.
func (n *Node) ResourcesUsageUpdateToParent(
	ctx context.Context,
) (host string, sessions uint, resourcesUsageNodesMap map[string]ResourcesUsage, err error) {
	return n.cmd.ResourcesUsageUpdateToParent(ctx)
}

// Get the update from the child nodes.
func (n *Node) ResourcesUsageUpdateFromChild(
	ctx context.Context,
	sessions uint,
	resourcesUsageNodesMap map[string]ResourcesUsage,
) (err error) {
	return n.cmd.ResourcesUsageUpdateFromChild(ctx, sessions, resourcesUsageNodesMap)
}

// Redirect new requests to the best offload target.
func (n *Node) RedirectNewRequests(
	ctx context.Context,
) (redirect bool, host string) {
	return n.cmd.RedirectNewRequests(ctx)
}
