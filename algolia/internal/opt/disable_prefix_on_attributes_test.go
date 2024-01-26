// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

func TestDisablePrefixOnAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.DisablePrefixOnAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DisablePrefixOnAttributes([]string{}...),
		},
		{
			opts:     []interface{}{opt.DisablePrefixOnAttributes("value1")},
			expected: opt.DisablePrefixOnAttributes("value1"),
		},
		{
			opts:     []interface{}{opt.DisablePrefixOnAttributes("value1", "value2", "value3")},
			expected: opt.DisablePrefixOnAttributes("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractDisablePrefixOnAttributes(c.opts...)
			out opt.DisablePrefixOnAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestDisablePrefixOnAttributes_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.DisablePrefixOnAttributesOption
	}{
		{
			payload:  `""`,
			expected: opt.DisablePrefixOnAttributes([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.DisablePrefixOnAttributes("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.DisablePrefixOnAttributes("value1", "value2", "value3"),
		},
	} {
		var got opt.DisablePrefixOnAttributesOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}
