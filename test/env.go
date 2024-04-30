package env_test

import (
	"context"
	"io"

	"github.com/ermes-labs/api-go/api"
)

var n1 api.Node
var n2 api.Node

func init() {
	sessionToken, err := n1.CreateSession(
		context.Background(),
		api.NewCreateSessionOptionsBuilder().Build())

	if err != nil {
		panic(err)
	}

	offloadedTo, err := n1.AcquireSession(
		context.Background(),
		sessionToken,
		api.NewAcquireSessionOptionsBuilder().Build(),
		func() error {
			// Do something...
			return nil
		})

	if err != nil {
		panic(err)
	}

	if offloadedTo != nil {
		// The session was offloaded to another node.
		// Do something...
		return
	}

	location, err := n1.OffloadSession(
		context.Background(),
		"session1",
		api.NewOffloadSessionOptionsBuilder().Build(),
		func(ctx context.Context, sm api.SessionMetadata, r io.Reader) (api.SessionLocation, error) {
			return n2.OnloadSession(ctx, sm, r, api.OnloadSessionOptions{})
		},
		func(ctx context.Context, oldLocation api.SessionLocation, newLocation api.SessionLocation) (bool, error) {
			var node *api.Node // = find the node from oldLocation.Host
			return node.UpdateOffloadedSessionLocation(ctx, oldLocation.SessionId, newLocation)
		})

	if err != nil {
		panic(err)
	}

	println(location.Host)
}
