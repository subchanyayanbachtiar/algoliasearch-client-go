// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractRestrictHighlightAndSnippetArrays returns the first found RestrictHighlightAndSnippetArraysOption from the
// given variadic arguments or nil otherwise.
func ExtractRestrictHighlightAndSnippetArrays(opts ...interface{}) *opt.RestrictHighlightAndSnippetArraysOption {
	for _, o := range opts {
		if v, ok := o.(*opt.RestrictHighlightAndSnippetArraysOption); ok {
			return v
		}
	}
	return nil
}
