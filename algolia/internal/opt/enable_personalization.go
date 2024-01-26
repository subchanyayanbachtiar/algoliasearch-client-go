// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractEnablePersonalization returns the first found EnablePersonalizationOption from the
// given variadic arguments or nil otherwise.
func ExtractEnablePersonalization(opts ...interface{}) *opt.EnablePersonalizationOption {
	for _, o := range opts {
		if v, ok := o.(*opt.EnablePersonalizationOption); ok {
			return v
		}
	}
	return nil
}
