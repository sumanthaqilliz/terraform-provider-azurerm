package targets

import "github.com/Azure/go-autorest/autorest"

type TargetsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTargetsClientWithBaseURI(endpoint string) TargetsClient {
	return TargetsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
