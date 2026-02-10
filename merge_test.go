package jsongo

import (
	"encoding/json"
	"strings"
	"testing"
)

func assertPanics(t *testing.T, expectedPanic string, f func()) {
	defer func() {
		if r := recover(); r != nil {
			if strings.Contains(r.(string), expectedPanic) {
				return
			}
			t.Errorf("Expected panic: %s, got: %v", expectedPanic, r)
		} else {
			t.Errorf("Expected panic: %s, got no panic", expectedPanic)
		}
	}()
	f()
}

type testMerge struct {
	name         string
	a            string
	b            string
	expectedA    string
	panicMessage string
}

func TestMerge(t *testing.T) {
	tests := []testMerge{
		{
			name:         "Not same type",
			a:            `{"d": [3, 2, 1],"b": {"c": 1}, "a": {}}`,
			b:            `{"d": [3, 2, 1, 4],"b": {"d": 2}, "a": []}`,
			expectedA:    `{"b": {"c": 1}, "a": {}}`,
			panicMessage: "cannot merge nodes of different types",
		},
		{
			name:         "Not same type 2",
			a:            `{"d": [3, 2, 1],"b": {"c": 1}, "a": {}}`,
			b:            `{"d": [3, 2, 1, 4],"b": {"d": 2}, "a": 1}`,
			expectedA:    `{"b": {"c": 1}, "a": {}}`,
			panicMessage: "cannot merge nodes of different types",
		},
		{
			name:         "Not same type 3",
			a:            `{"d": [3, 2, 1],"b": {"c": 1}, "a": []}`,
			b:            `{"d": [3, 2, 1, 4],"b": {"d": 2}, "a": 1}`,
			expectedA:    `{"b": {"c": 1}, "a": {}}`,
			panicMessage: "cannot merge nodes of different types",
		},
		{
			name:         "Value already set and different",
			a:            `{"d": [3, 2, 1],"b": {"c": 1}, "a": 1}`,
			b:            `{"d": [3, 2, 1, 4],"b": {"d": 2}, "a": 2}`,
			expectedA:    `{"b": {"c": 1}, "a": {}}`,
			panicMessage: "value already set",
		},
		{
			name:         "Value already set and same",
			a:            `{"d": [3, 2, 1],"b": {"c": 1}, "a": 1}`,
			b:            `{"d": [3, 2, 1, 4], "b": {"d": 2}, "a": 1}`,
			expectedA:    `{"a":1,"b":{"c":1,"d":2},"d":[3,2,1,4]}`,
			panicMessage: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			aNode := &Node{}
			bNode := &Node{}
			if err := aNode.UnmarshalJSON([]byte(test.a)); err != nil {
				t.Errorf("Failed to unmarshal a: %v", err)
			}
			if err := bNode.UnmarshalJSON([]byte(test.b)); err != nil {
				t.Errorf("Failed to unmarshal b: %v", err)
			}
			if test.panicMessage != "" {
				assertPanics(t, test.panicMessage, func() {
					aNode.Merge(bNode)
				})
			} else {
				aNode.Merge(bNode)
				jsonA, err := json.Marshal(aNode)
				if err != nil {
					t.Errorf("Failed to marshal a: %v", err)
				}
				if string(jsonA) != test.expectedA {
					t.Errorf("Expected a to be %s, got %s", test.expectedA, string(jsonA))
				}
			}
		})
	}
}
