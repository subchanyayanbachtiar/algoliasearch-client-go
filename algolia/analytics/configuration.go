package analytics

import (
	"time"

	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/region"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/transport"
)

// Configuration contains all the different parameters one can change to
// instantiate a new client for the Analytics API.
type Configuration struct {
	AppID          string
	APIKey         string
	Hosts          []string
	Requester      transport.Requester
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	Region         region.Region
	Headers        map[string]string
	ExtraUserAgent string
}
