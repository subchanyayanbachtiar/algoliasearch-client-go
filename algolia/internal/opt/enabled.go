// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractEnabled returns the first found EnabledOption from the
// given variadic arguments or nil otherwise.
func ExtractEnabled(opts ...interface{}) *opt.EnabledOption {
	for _, o := range opts {
		if v, ok := o.(*opt.EnabledOption); ok {
			return v
		}
	}
	return nil
}
