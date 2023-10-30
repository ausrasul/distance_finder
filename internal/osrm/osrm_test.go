package osrm

import (
	"reflect"
	"testing"

	"example.com/dist_finder/internal/api"
)

type MockApi struct {
	ExpectedReturn api.OsrmResp
	ExpectedErr    error
}

func (mockApi MockApi) GetRoutes(coords []string) (api.OsrmResp, error) {
	return mockApi.ExpectedReturn, mockApi.ExpectedErr
}

func Test_GetRoutes(t *testing.T) {
	cases := []struct {
		desc        string
		apiData     api.OsrmResp
		apiErr      error
		inputCoords []string
		expectedOk  bool
		expected    Reply
	}{
		{
			desc: "Create reply and return ok",
			apiData: api.OsrmResp{
				Code:      "Ok",
				Distances: [][]float32{{0.0, 1.1, 2.2}},
				Durations: [][]float32{{0.0, 3.3, 4.4}},
			},
			apiErr: nil,
			inputCoords: []string{
				"aaa",
				"bbb",
				"ccc",
			},
			expectedOk: true,
			expected: Reply{
				Source: "aaa",
				Routes: []Route{
					{
						Destination: "bbb",
						Distance:    1.1,
						Duration:    3.3,
					},
					{
						Destination: "ccc",
						Distance:    2.2,
						Duration:    4.4,
					},
				},
			},
		},
		{
			desc: "Create reply and return ok",
			apiData: api.OsrmResp{
				Code:      "Ok",
				Distances: [][]float32{{0.0, 1.1}},
				Durations: [][]float32{{0.0, 3.3}},
			},
			inputCoords: []string{
				"aaa",
				"bbb",
			},
			expectedOk: true,
			expected: Reply{
				Source: "aaa",
				Routes: []Route{
					{
						Destination: "bbb",
						Distance:    1.1,
						Duration:    3.3,
					},
				},
			},
		},
		{
			desc: "Sort routes by Duration and Distance (if Duration is equal)",
			apiData: api.OsrmResp{
				Code:      "Ok",
				Distances: [][]float32{{0.0, 1.1, 4.4, 1.1, 3.3, 2.2, 5.5, 6.6, 2.2, 9.9}},
				Durations: [][]float32{{0.0, 3.3, 4.4, 4.4, 4.4, 4.4, 1.1, 2.2, 1.0, 5.5}},
			},
			inputCoords: []string{
				"aaa",
				"bbb",
				"cc4",
				"cc1",
				"cc3",
				"cc2",
				"ddd",
				"eee",
				"fff",
				"ggg",
			},
			expectedOk: true,
			expected: Reply{
				Source: "aaa",
				Routes: []Route{
					{
						Destination: "fff",
						Distance:    2.2,
						Duration:    1.0,
					},
					{
						Destination: "ddd",
						Distance:    5.5,
						Duration:    1.1,
					},
					{
						Destination: "eee",
						Distance:    6.6,
						Duration:    2.2,
					},
					{
						Destination: "bbb",
						Distance:    1.1,
						Duration:    3.3,
					},
					{
						Destination: "cc1",
						Distance:    1.1,
						Duration:    4.4,
					},
					{
						Destination: "cc2",
						Distance:    2.2,
						Duration:    4.4,
					},
					{
						Destination: "cc3",
						Distance:    3.3,
						Duration:    4.4,
					},
					{
						Destination: "cc4",
						Distance:    4.4,
						Duration:    4.4,
					},
					{
						Destination: "ggg",
						Distance:    9.9,
						Duration:    5.5,
					},
				},
			},
		},
		{
			desc: "Don't create reply if no distance",
			apiData: api.OsrmResp{
				Code:      "Ok",
				Distances: [][]float32{},
				Durations: [][]float32{{0.0, 3.3, 4.4}},
			},
			inputCoords: []string{
				"aaa",
				"bbb",
				"ccc",
			},
			expectedOk: false,
			expected:   Reply{},
		},
		{
			desc: "Don't create reply if no distance",
			apiData: api.OsrmResp{
				Code:      "Ok",
				Distances: [][]float32{{}},
				Durations: [][]float32{{0.0, 3.3, 4.4}},
			},
			inputCoords: []string{
				"aaa",
				"bbb",
				"ccc",
			},
			expectedOk: false,
			expected:   Reply{},
		},
		{
			desc: "Don't create reply if no durations",
			apiData: api.OsrmResp{
				Code:      "Ok",
				Distances: [][]float32{{0.0, 3.3, 4.4}},
				Durations: [][]float32{{}},
			},
			inputCoords: []string{
				"aaa",
				"bbb",
				"ccc",
			},
			expectedOk: false,
			expected:   Reply{},
		},
		{
			desc: "Don't create reply when return values not same number as request's",
			apiData: api.OsrmResp{
				Code:      "Ok",
				Distances: [][]float32{{0.0}},
				Durations: [][]float32{{0.0, 3.3}},
			},
			inputCoords: []string{
				"aaa",
				"bbb",
			},
			expectedOk: false,
			expected:   Reply{},
		},
		{
			desc: "Don't create reply when return values not same number as request's",
			apiData: api.OsrmResp{
				Code:      "Ok",
				Distances: [][]float32{{0.0, 1, 1}},
				Durations: [][]float32{{0.0}},
			},
			inputCoords: []string{
				"aaa",
				"bbb",
			},
			expectedOk: false,
			expected:   Reply{},
		},
		{
			desc: "Don't create reply if response has code other than Ok",
			apiData: api.OsrmResp{
				Code:      "NOK",
				Distances: [][]float32{{0.0, 1.1}},
				Durations: [][]float32{{0.0, 2.2}},
			},
			inputCoords: []string{
				"aaa",
				"bbb",
			},
			expectedOk: false,
			expected:   Reply{},
		},
	}
	for _, testCase := range cases {
		osrm := OSRM{
			api: &MockApi{
				ExpectedReturn: testCase.apiData,
				ExpectedErr:    testCase.apiErr,
			},
		}
		reply, ok := osrm.GetRoutes(testCase.inputCoords)
		if ok != testCase.expectedOk {
			t.Error("Expected ", testCase.expectedOk, " got ", ok)
		}
		if !reflect.DeepEqual(testCase.expected, reply) {
			t.Error("Expected ", testCase.expected, " got ", reply)
		}
	}
}
