// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractDecompoundedAttributes returns the first found DecompoundedAttributesOption from the
// given variadic arguments or nil otherwise.
func ExtractDecompoundedAttributes(opts ...interface{}) *opt.DecompoundedAttributesOption {
	for _, o := range opts {
		if v, ok := o.(*opt.DecompoundedAttributesOption); ok {
			return v
		}
	}
	return nil
}
