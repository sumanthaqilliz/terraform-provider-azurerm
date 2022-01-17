package chaos

import (
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

var _ sdk.UntypedServiceRegistration = Registration{}

type Registration struct {
}

func (r Registration) Name() string {
	return "Chaos Studio"
}

func (r Registration) WebsiteCategories() []string {
	return []string{
		"Chaos Studio",
	}
}

func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{}
}

func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{}
}
