package api

import (
	"context"

	"github.com/ermes-labs/api-go/infrastructure"
)

type ResourcesUsage = map[string]float64
type ResourcesUsageIndex = map[string]float64

// Commands to get and update the resources usage of the sessions and the nodes.
type ResourcesUsageCommands interface {
	// Load the infrastructure.
	LoadInfrastructure(
		ctx context.Context,
		infrastructure infrastructure.Infrastructure,
	) (err error)
	// Get the parent node of a node.
	GetParentNodeOf(
		ctx context.Context,
		nodeId string,
	) (*infrastructure.Node, error)
	// Get the children nodes of a node.
	GetChildrenNodesOf(
		ctx context.Context,
		nodeId string,
	) ([]infrastructure.Node, error)
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
	) (node infrastructure.Node, sessions uint, resourcesUsageNodesMap map[string]ResourcesUsage, err error)
	// Get the update from the child nodes.
	ResourcesUsageUpdateFromChild(
		ctx context.Context,
		sessions uint,
		resourcesUsageNodesMap map[string]ResourcesUsage,
	) (err error)
}

// load the infrastructure.
func (n *Node) LoadInfrastructure(
	ctx context.Context,
	infrastructure infrastructure.Infrastructure,
) (err error) {
	return n.Cmd.LoadInfrastructure(ctx, infrastructure)
}

// Get the parent node of a node.
func (n *Node) GetParentNodeOf(
	ctx context.Context,
	nodeId string,
) (*infrastructure.Node, error) {
	return n.Cmd.GetParentNodeOf(ctx, nodeId)
}

// Get the children nodes of a node.
func (n *Node) GetChildrenNodesOf(
	ctx context.Context,
	nodeId string,
) ([]infrastructure.Node, error) {
	return n.Cmd.GetChildrenNodesOf(ctx, nodeId)
}

// Get the resources usage of a session.
// errors:
// - ErrSessionNotFound: If no session with the given id is found.
func (n *Node) GetSessionResourcesUsage(
	ctx context.Context,
	sessionId string,
) (resourcesUsage ResourcesUsage, err error) {
	return n.Cmd.GetSessionResourcesUsage(ctx, sessionId)
}

// Get the resources usage of a node.
func (n *Node) GetNodeResourcesUsage(
	ctx context.Context,
	nodeId string,
) (sessions uint, resourcesUsage ResourcesUsage, err error) {
	return n.Cmd.GetNodeResourcesUsage(ctx, nodeId)
}

// Get the resources usage of all the nodes.
func (n *Node) GetNodeResourcesUsageIndex(
	ctx context.Context,
	nodeId string,
) (resourcesUsageIndex ResourcesUsageIndex, err error) {
	_, resourcesUsage, err := n.Cmd.GetNodeResourcesUsage(ctx, nodeId)

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
	return n.Cmd.UpdateSessionResourcesUsage(ctx, sessionId, resourcesUsage)
}

// Get the update to send to the parent node.
func (n *Node) ResourcesUsageUpdateToParent(
	ctx context.Context,
) (node infrastructure.Node, sessions uint, resourcesUsageNodesMap map[string]ResourcesUsage, err error) {
	return n.Cmd.ResourcesUsageUpdateToParent(ctx)
}

// Get the update from the child nodes.
func (n *Node) ResourcesUsageUpdateFromChild(
	ctx context.Context,
	sessions uint,
	resourcesUsageNodesMap map[string]ResourcesUsage,
) (err error) {
	return n.Cmd.ResourcesUsageUpdateFromChild(ctx, sessions, resourcesUsageNodesMap)
}
