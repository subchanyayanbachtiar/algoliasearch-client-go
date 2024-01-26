// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractType returns the first found TypeOption from the
// given variadic arguments or nil otherwise.
func ExtractType(opts ...interface{}) *opt.TypeOption {
	for _, o := range opts {
		if v, ok := o.(*opt.TypeOption); ok {
			return v
		}
	}
	return nil
}
