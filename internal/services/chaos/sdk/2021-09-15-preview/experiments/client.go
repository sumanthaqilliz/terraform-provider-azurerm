package experiments

import "github.com/Azure/go-autorest/autorest"

type ExperimentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewExperimentsClientWithBaseURI(endpoint string) ExperimentsClient {
	return ExperimentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
