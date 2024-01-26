// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractValidUntil returns the first found ValidUntilOption from the
// given variadic arguments or nil otherwise.
func ExtractValidUntil(opts ...interface{}) *opt.ValidUntilOption {
	for _, o := range opts {
		if v, ok := o.(*opt.ValidUntilOption); ok {
			return v
		}
	}
	return nil
}
