package api

import "github.com/ermes-labs/api-go/infrastructure"

type Node struct {
	Cmd Commands
	infrastructure.Node
}

func NewNode(node infrastructure.Node, cmd Commands) *Node {
	return &Node{
		Cmd:  cmd,
		Node: node,
	}
}

type Commands interface {
	AcquireSessionCommands
	BestOffloadTargetsCommands
	CreateAndAcquireSessionCommands
	CreateSessionCommands
	GarbageCollectSessionsCommands
	OffloadSessionCommands
	OnloadSessionCommands
	ResourcesUsageCommands
	SessionMetadataCommands
}
