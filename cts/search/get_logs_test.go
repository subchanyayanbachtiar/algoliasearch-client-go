package search

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/cts"
)

func TestGetLogs(t *testing.T) {
	t.Parallel()
	c := cts.InitSearchClient1(t)

	for i := 0; i < 2; i++ {
		_, err := c.ListIndices()
		require.NoError(t, err)
	}

	res, err := c.GetLogs(
		opt.Offset(0),
		opt.Length(2),
		opt.Type("all"),
	)
	require.NoError(t, err)
	require.Len(t, res.Logs, 2)
}
