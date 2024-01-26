// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractRestrictSources returns the first found RestrictSourcesOption from the
// given variadic arguments or nil otherwise.
func ExtractRestrictSources(opts ...interface{}) *opt.RestrictSourcesOption {
	for _, o := range opts {
		if v, ok := o.(*opt.RestrictSourcesOption); ok {
			return v
		}
	}
	return nil
}
