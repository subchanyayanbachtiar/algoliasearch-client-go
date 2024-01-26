package index

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/search"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/cts"
)

func TestExists(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	ok, err := index.Exists()
	require.False(t, ok)
	require.NoError(t, err)

	res, err := index.SaveObject(map[string]string{"attribute": "value"})
	require.NoError(t, err)
	err = res.Wait()
	require.NoError(t, err)

	ok, _ = index.Exists()
	require.True(t, ok)
	require.NoError(t, err)

	ok, err = search.NewClient(index.GetAppID(), "invalid").InitIndex(index.GetName()).Exists()
	require.False(t, ok)
	require.Error(t, err)
}
