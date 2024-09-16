package json_test

import (
	"testing"

	"github.com/hlcfan/fm/json"
	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	tcs := []struct {
		description string
		input       string
		tokens      []json.Token
	}{
		{
			description: "It lexers json",
			input:       `{"abccc":  3, "bbb": true, "cc":  false, "dd": [{"d1":1}, {"d2": null}]}`,
			tokens: []json.Token{
				{Kind: 7, Value: "{"},
				{Kind: 3, Value: `"abccc"`},
				{Kind: 4, Value: ":"},
				{Kind: 2, Value: "3"},
				{Kind: 4, Value: ","},
				{Kind: 3, Value: `"bbb"`},
				{Kind: 4, Value: ":"},
				{Kind: 2, Value: "true"},
				{Kind: 4, Value: ","},
				{Kind: 3, Value: `"cc"`},
				{Kind: 4, Value: ":"},
				{Kind: 0, Value: "false"},
				{Kind: 4, Value: ","},
				{Kind: 3, Value: `"dd"`},
				{Kind: 4, Value: ":"},
				{Kind: 5, Value: "["},
				{Kind: 7, Value: "{"},
				{Kind: 3, Value: `"d1"`},
				{Kind: 4, Value: ":"},
				{Kind: 2, Value: "1"},
				{Kind: 8, Value: "}"},
				{Kind: 4, Value: ","},
				{Kind: 7, Value: "{"},
				{Kind: 3, Value: `"d2"`},
				{Kind: 4, Value: ":"},
				{Kind: 1, Value: "null"},
				{Kind: 8, Value: "}"},
				{Kind: 6, Value: "]"},
				{Kind: 8, Value: "}"},
			},
		},
		{
			description: "It lexers incomplete string value in json",
			input:       `{"abccc":  "what-the`,
			tokens: []json.Token{
				{Kind: 7, Value: "{"},
				{Kind: 3, Value: `"abccc"`},
				{Kind: 4, Value: ":"},
				{Kind: 3, Value: "\"what-the"},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			scanner := json.NewScanner(tc.input)
			tokens := scanner.Scan()
			assert.Equal(t, tc.tokens, tokens)
		})
	}
}
