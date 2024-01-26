package search

import (
	iopt "github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

type KeyQueryParams struct {
	RestrictSources *opt.RestrictSourcesOption `json:"restrictSources,omitempty"`
	QueryParams
}

func newKeyQueryParams(opts ...interface{}) KeyQueryParams {
	return KeyQueryParams{
		RestrictSources: iopt.ExtractRestrictSources(opts...),
		QueryParams:     newQueryParams(opts...),
	}
}
