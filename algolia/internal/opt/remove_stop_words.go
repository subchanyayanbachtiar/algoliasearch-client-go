// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractRemoveStopWords returns the first found RemoveStopWordsOption from the
// given variadic arguments or nil otherwise.
func ExtractRemoveStopWords(opts ...interface{}) *opt.RemoveStopWordsOption {
	for _, o := range opts {
		if v, ok := o.(*opt.RemoveStopWordsOption); ok {
			return v
		}
	}
	return nil
}
