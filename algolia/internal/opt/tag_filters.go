// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractTagFilters returns the first found TagFiltersOption from the
// given variadic arguments or nil otherwise.
func ExtractTagFilters(opts ...interface{}) *opt.TagFiltersOption {
	for _, o := range opts {
		if v, ok := o.(*opt.TagFiltersOption); ok {
			return v
		}
	}
	return nil
}
