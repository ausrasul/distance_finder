package api

import (
	"testing"
)

func Test_CreateOSRMReqUrl(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{
			input: []string{
				"11.11,22.22",
				"33.33,44.44",
				"55.55,66.66",
			},
			expected: "http://router.project-osrm.org/table/v1/driving/11.11,22.22;33.33,44.44;55.55,66.66?sources=0&annotations=distance,duration",
		},
		{
			input: []string{
				"77.77,11.11",
				"55.55,66.66",
				"33.33,44.44",
			},
			expected: "http://router.project-osrm.org/table/v1/driving/77.77,11.11;55.55,66.66;33.33,44.44?sources=0&annotations=distance,duration",
		},
		{
			input: []string{
				"11.11,22.22",
				"33.33,44.44",
			},
			expected: "http://router.project-osrm.org/table/v1/driving/11.11,22.22;33.33,44.44?sources=0&annotations=distance,duration",
		},
		{
			input: []string{
				"11.11,22.22",
				"33.33,44.44",
				"33.33,44.44",
				"33.33,44.44",
			},
			expected: "http://router.project-osrm.org/table/v1/driving/11.11,22.22;33.33,44.44;33.33,44.44;33.33,44.44?sources=0&annotations=distance,duration",
		},
	}
	for _, testCase := range cases {
		osrcReq := createUrl(testCase.input)
		if osrcReq != testCase.expected {
			t.Error("Expected ", testCase.expected, " got ", osrcReq)
		}
	}
}
