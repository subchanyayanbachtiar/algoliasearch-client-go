// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractTypoTolerance returns the first found TypoToleranceOption from the
// given variadic arguments or nil otherwise.
func ExtractTypoTolerance(opts ...interface{}) *opt.TypoToleranceOption {
	for _, o := range opts {
		if v, ok := o.(*opt.TypoToleranceOption); ok {
			return v
		}
	}
	return nil
}
