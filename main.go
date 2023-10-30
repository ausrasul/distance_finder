package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"

	Osrm "example.com/dist_finder/internal/osrm"
)

func main() {
	go runWebServer("8080")
	// Here I would implement some telemetry
	// and may be another web server for interal queries like KPIs etc.
	select {}
}
func runWebServer(port string) {
	http.HandleFunc("/routes", handleRequests)

	log.Print("Listening on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("Recovered error while handling a request ", r, string(debug.Stack()))
		}
	}()

	coords, ok := parseQuery(r.URL.Query())
	if !ok {
		log.Print("Cannot parse url query")
		return
	}
	osrm := Osrm.New()
	reply, ok := osrm.GetRoutes(coords)
	if !ok {
		log.Print("error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	json.NewEncoder(w).Encode(reply)

}

func parseQuery(input map[string][]string) (coords []string, _ok bool) {
	src, ok := input["src"]
	if !ok || len(src) != 1 {
		return []string{}, false
	}
	dst, ok := input["dst"]
	if !ok || len(dst) == 0 {
		return []string{}, false
	}
	coords = append(coords, src[0])
	coords = append(coords, dst...)
	return coords, true
}
