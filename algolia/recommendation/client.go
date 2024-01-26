package recommendation

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/call"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/compression"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/region"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/transport"
)

// Deprecated: use personalization.Client instead
// Client provides methods to interact with the Algolia Recommendation API.
type Client struct {
	transport *transport.Transport
}

// Deprecated: use personalization.NewClient() instead
// NewClient instantiates a new client able to interact with the Algolia
// Recommendation API.
func NewClient(appID, apiKey string, region region.Region) *Client {
	return NewClientWithConfig(
		Configuration{
			AppID:  appID,
			APIKey: apiKey,
			Region: region,
		},
	)
}

// Deprecated: use personalization.NewClientWithConfig() instead
// NewClientWithConfig instantiates a new client able to interact with the
// Recommendation API.
func NewClientWithConfig(config Configuration) *Client {
	var hosts []*transport.StatefulHost

	if config.Hosts == nil {
		hosts = defaultHosts(config.Region)
	} else {
		for _, h := range config.Hosts {
			hosts = append(hosts, transport.NewStatefulHost(h, call.IsReadWrite))
		}
	}

	return &Client{
		transport: transport.New(
			hosts,
			config.Requester,
			config.AppID,
			config.APIKey,
			config.ReadTimeout,
			config.WriteTimeout,
			config.Headers,
			config.ExtraUserAgent,
			compression.None,
		),
	}
}
