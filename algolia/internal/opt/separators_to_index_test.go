// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

func TestSeparatorsToIndex(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.SeparatorsToIndexOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.SeparatorsToIndex(""),
		},
		{
			opts:     []interface{}{opt.SeparatorsToIndex("")},
			expected: opt.SeparatorsToIndex(""),
		},
		{
			opts:     []interface{}{opt.SeparatorsToIndex("content of the string value")},
			expected: opt.SeparatorsToIndex("content of the string value"),
		},
	} {
		var (
			in  = ExtractSeparatorsToIndex(c.opts...)
			out opt.SeparatorsToIndexOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
