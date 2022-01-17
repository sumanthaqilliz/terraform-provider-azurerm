package targettypes

import "github.com/Azure/go-autorest/autorest"

type TargetTypesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTargetTypesClientWithBaseURI(endpoint string) TargetTypesClient {
	return TargetTypesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
