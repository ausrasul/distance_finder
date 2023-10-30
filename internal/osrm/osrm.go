package osrm

import (
	"sort"

	"example.com/dist_finder/internal/api"
)

type Reply struct {
	Source string  `json:"source"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Destination string  `json:"destination"`
	Duration    float32 `json:"duration"`
	Distance    float32 `json:"distance"`
}

// used to facilitate testing by replacing the API with its http calls.
type apiInterface interface {
	GetRoutes(coords []string) (api.OsrmResp, error)
	// Add other package-specific methods here
}

type OSRM struct {
	api apiInterface
}

func New() OSRM {
	return OSRM{
		api: &api.Api{},
	}
}

func (osrm OSRM) GetRoutes(coords []string) (reply Reply, ok bool) {
	osrmResp, err := osrm.api.GetRoutes(coords)
	if err != nil {
		return
	}
	if len(osrmResp.Distances) <= 0 ||
		len(osrmResp.Distances[0]) <= 0 ||
		len(osrmResp.Durations[0]) <= 0 ||
		len(osrmResp.Distances[0]) != len(coords) ||
		len(osrmResp.Durations[0]) != len(coords) ||
		osrmResp.Code != "Ok" {
		return reply, false
	}
	reply.Source = coords[0]
	for i, dst := range coords[1:] {
		route := Route{
			Destination: dst,
			Distance:    osrmResp.Distances[0][i+1],
			Duration:    osrmResp.Durations[0][i+1],
		}
		reply.Routes = append(reply.Routes, route)
	}
	sort.Slice(reply.Routes, func(i, j int) bool {
		if reply.Routes[i].Duration == reply.Routes[j].Duration {
			return reply.Routes[i].Distance < reply.Routes[j].Distance
		} else {
			return reply.Routes[i].Duration < reply.Routes[j].Duration
		}
	})
	return reply, true
}
