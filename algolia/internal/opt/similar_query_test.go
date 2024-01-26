// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

func TestSimilarQuery(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.SimilarQueryOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.SimilarQuery(""),
		},
		{
			opts:     []interface{}{opt.SimilarQuery("")},
			expected: opt.SimilarQuery(""),
		},
		{
			opts:     []interface{}{opt.SimilarQuery("content of the string value")},
			expected: opt.SimilarQuery("content of the string value"),
		},
	} {
		var (
			in  = ExtractSimilarQuery(c.opts...)
			out opt.SimilarQueryOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
