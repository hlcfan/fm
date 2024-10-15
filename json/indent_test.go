package json_test

import (
	"testing"

	"github.com/hlcfan/fm/json"
	"github.com/stretchr/testify/assert"
)

func TestIndent(t *testing.T) {
	tcs := []struct {
		description string
		// tokens      []json.Token
		input  string
		output string
	}{
		{
			description: "It indents empty object",
			input:       `{}`,
			output: `{
}`,
		},
		{
			description: "It indents empty array",
			input:       `[]`,
			output: `[
]`,
		},
		{
			description: "It indents array",
			input:       `[1,2,3]`,
			output: `[
  1,
  2,
  3
]`,
		},
		{
			description: "It indents complex object",
			input:       `{"a":1,"bool":false,"string":"foo","array":[1,2,3]}`,
			output: `{
  "a": 1,
  "bool": false,
  "string": "foo",
  "array": [
    1,
    2,
    3
  ]
}`,
		},
		{
			description: "It indents complex object",
			input:       `{"abccc":  3, "bbb": {}, "cc":  false, "dd": [{"d1":1.30}, {"d2": null, "d3": "1", "d4": -3.14159265358}], "foo":"bar","char":"a\a\"\n"}`,
			output: `{
  "abccc": 3,
  "bbb": {
  },
  "cc": false,
  "dd": [
    {
      "d1": 1.30
    },
    {
      "d2": null,
      "d3": "1",
      "d4": -3.14159265358
    }
  ],
  "foo": "bar"
}`,
		},
	}

	for _, tc := range tcs {
		defaultIndent := "  "
		t.Run(tc.description, func(t *testing.T) {
			s := json.NewScanner(tc.input)
			output := json.Indent(s.Scan(), defaultIndent)
			assert.Equal(t, tc.output, output.String())
		})
	}
}
