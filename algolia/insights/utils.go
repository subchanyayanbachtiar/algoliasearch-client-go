package insights

import (
	"fmt"

	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/call"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/region"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/transport"
)

func defaultHosts(r region.Region) (hosts []*transport.StatefulHost) {
	switch r {
	case region.EU, region.US:
		hosts = append(hosts, transport.NewStatefulHost(fmt.Sprintf("insights.%s.algolia.io", r), call.IsReadWrite))
	default:
		hosts = append(hosts, transport.NewStatefulHost("insights.algolia.io", call.IsReadWrite))
	}
	return
}
