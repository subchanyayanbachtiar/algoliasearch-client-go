// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractRanking returns the first found RankingOption from the
// given variadic arguments or nil otherwise.
func ExtractRanking(opts ...interface{}) *opt.RankingOption {
	for _, o := range opts {
		if v, ok := o.(*opt.RankingOption); ok {
			return v
		}
	}
	return nil
}
