package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type OsrmResp struct {
	Code      string      `json:"code"`
	Distances [][]float32 `json:"distances"`
	Durations [][]float32 `json:"durations"`
}

type Api struct{}

func (Api) GetRoutes(coords []string) (OsrmResp, error) {
	url := createUrl(coords)
	res, err := http.Get(url)
	if err != nil {
		return OsrmResp{}, err
	}
	if res.StatusCode != 200 {
		return OsrmResp{}, errors.New(fmt.Sprintf("OSRM http return status %d", res.StatusCode))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return OsrmResp{}, err
	}
	osrmResp := OsrmResp{}
	err = json.Unmarshal(resBody, &osrmResp)
	if err != nil {
		return OsrmResp{}, err
	}
	return osrmResp, nil
}

func createUrl(coords []string) string {
	/*
			I used Table service call and not Route service that is provided in the assignment.
		   This seem to fullfill the requirement with a single call.
	*/
	url := "http://router.project-osrm.org/table/v1/driving/"
	for _, coord := range coords {
		url += coord + ";"
	}
	url = url[:len(url)-1]
	url += "?sources=0&annotations=distance,duration"
	return url
}
