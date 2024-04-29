package triggered_functions

import (
	"context"

	"github.com/ermes-labs/api-go/api"
	http_functions "github.com/ermes-labs/api-go/functions/http"
)

func SendStatus(
	n *api.Node,
) error {
	// Extract from the body the
	_, sessions, resourcesUsageMap, err := n.ResourcesUsageUpdateToParent(context.Background())
	if err != nil {
		return err
	}

	var handler http_functions.Handler
	err = handler.IssueReceiveStatusRequest(context.Background(), sessions, resourcesUsageMap)

	return err
}
