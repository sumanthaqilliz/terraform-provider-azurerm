package capabilities

import "github.com/Azure/go-autorest/autorest"

type CapabilitiesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCapabilitiesClientWithBaseURI(endpoint string) CapabilitiesClient {
	return CapabilitiesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
