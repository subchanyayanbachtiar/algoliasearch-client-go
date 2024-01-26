// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

func TestFilterPromotes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.FilterPromotesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.FilterPromotes(false),
		},
		{
			opts:     []interface{}{opt.FilterPromotes(true)},
			expected: opt.FilterPromotes(true),
		},
		{
			opts:     []interface{}{opt.FilterPromotes(false)},
			expected: opt.FilterPromotes(false),
		},
	} {
		var (
			in  = ExtractFilterPromotes(c.opts...)
			out opt.FilterPromotesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
