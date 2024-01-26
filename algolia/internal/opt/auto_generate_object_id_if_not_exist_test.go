// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

func TestAutoGenerateObjectIDIfNotExist(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AutoGenerateObjectIDIfNotExistOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AutoGenerateObjectIDIfNotExist(false),
		},
		{
			opts:     []interface{}{opt.AutoGenerateObjectIDIfNotExist(true)},
			expected: opt.AutoGenerateObjectIDIfNotExist(true),
		},
		{
			opts:     []interface{}{opt.AutoGenerateObjectIDIfNotExist(false)},
			expected: opt.AutoGenerateObjectIDIfNotExist(false),
		},
	} {
		var (
			in  = ExtractAutoGenerateObjectIDIfNotExist(c.opts...)
			out opt.AutoGenerateObjectIDIfNotExistOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
