package capabilitytypes

import "github.com/Azure/go-autorest/autorest"

type CapabilityTypesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCapabilityTypesClientWithBaseURI(endpoint string) CapabilityTypesClient {
	return CapabilityTypesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
