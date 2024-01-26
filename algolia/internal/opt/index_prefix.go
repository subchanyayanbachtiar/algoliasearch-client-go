// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractIndexPrefix returns the first found IndexPrefixOption from the
// given variadic arguments or nil otherwise.
func ExtractIndexPrefix(opts ...interface{}) *opt.IndexPrefixOption {
	for _, o := range opts {
		if v, ok := o.(*opt.IndexPrefixOption); ok {
			return v
		}
	}
	return nil
}
