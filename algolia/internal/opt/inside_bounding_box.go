// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractInsideBoundingBox returns the first found InsideBoundingBoxOption from the
// given variadic arguments or nil otherwise.
func ExtractInsideBoundingBox(opts ...interface{}) *opt.InsideBoundingBoxOption {
	for _, o := range opts {
		if v, ok := o.(*opt.InsideBoundingBoxOption); ok {
			return v
		}
	}
	return nil
}
