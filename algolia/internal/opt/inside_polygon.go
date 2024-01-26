// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractInsidePolygon returns the first found InsidePolygonOption from the
// given variadic arguments or nil otherwise.
func ExtractInsidePolygon(opts ...interface{}) *opt.InsidePolygonOption {
	for _, o := range opts {
		if v, ok := o.(*opt.InsidePolygonOption); ok {
			return v
		}
	}
	return nil
}
