package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

func TestRemoveStopWords(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.RemoveStopWordsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.RemoveStopWords(false),
		},
		{
			opts:     []interface{}{opt.RemoveStopWords(true)},
			expected: opt.RemoveStopWords(true),
		},
		{
			opts:     []interface{}{opt.RemoveStopWords(false)},
			expected: opt.RemoveStopWords(false),
		},
		{
			opts:     []interface{}{opt.RemoveStopWordsFor("fr", "en")},
			expected: opt.RemoveStopWordsFor("fr", "en"),
		},
	} {
		var (
			in  = ExtractRemoveStopWords(c.opts...)
			out opt.RemoveStopWordsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
