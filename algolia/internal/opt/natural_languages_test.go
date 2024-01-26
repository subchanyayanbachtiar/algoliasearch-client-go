// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

func TestNaturalLanguages(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.NaturalLanguagesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.NaturalLanguages([]string{}...),
		},
		{
			opts:     []interface{}{opt.NaturalLanguages("value1")},
			expected: opt.NaturalLanguages("value1"),
		},
		{
			opts:     []interface{}{opt.NaturalLanguages("value1", "value2", "value3")},
			expected: opt.NaturalLanguages("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractNaturalLanguages(c.opts...)
			out opt.NaturalLanguagesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestNaturalLanguages_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.NaturalLanguagesOption
	}{
		{
			payload:  `""`,
			expected: opt.NaturalLanguages([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.NaturalLanguages("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.NaturalLanguages("value1", "value2", "value3"),
		},
	} {
		var got opt.NaturalLanguagesOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}
