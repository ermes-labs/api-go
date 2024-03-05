package api

import "github.com/ermes-labs/api-go/infrastructure"

type Node struct {
	cmd Commands
	infrastructure.Node
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
