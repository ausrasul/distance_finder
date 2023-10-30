package main

import (
	"reflect"
	"testing"
)

func Test_ParseQuery(t *testing.T) {
	cases := []struct {
		desc        string
		input       map[string][]string
		expectedOk  bool
		expectedVal []string
	}{
		{
			desc: "Parse correctly",
			input: map[string][]string{
				"src": {"12.23,12.32"},
				"dst": {"12.12,12.12", "23.23,23.23"},
			},
			expectedOk: true,
			expectedVal: []string{
				"12.23,12.32",
				"12.12,12.12",
				"23.23,23.23",
			},
		},
		{
			desc: "Don't parse if no source",
			input: map[string][]string{
				"dst": {"12.12,12.12", "23.23,23.23"},
			},
			expectedOk:  false,
			expectedVal: []string{},
		},
		{
			desc: "Don't parse if no destination",
			input: map[string][]string{
				"src": {"12.23,12.32"},
			},
			expectedOk:  false,
			expectedVal: []string{},
		},
		{
			desc: "Don't parse if multiple sources",
			input: map[string][]string{
				"src": {"12.12,12.12", "23.23,23.23"},
				"dst": {"12.12,12.12", "23.23,23.23"},
			},
			expectedOk:  false,
			expectedVal: []string{},
		},
		{
			desc: "Don't parse if multiple sources",
			input: map[string][]string{
				"src": {"12.12,12.12", "23.23,23.23"},
				"dst": {"12.12,12.12", "23.23,23.23"},
			},
			expectedOk:  false,
			expectedVal: []string{},
		},
		/*
			I can also regex check each coordinate, but this time I'll let the 3rd party API handle that instead.
		*/
	}
	for _, testCase := range cases {
		parsed, ok := parseQuery(testCase.input)
		if ok != testCase.expectedOk {
			t.Error("Expected ", testCase.expectedOk, " got ", ok)
		}
		if !reflect.DeepEqual(parsed, testCase.expectedVal) {
			t.Error("Expected ", testCase.expectedVal, " got ", parsed)
		}
	}
}
