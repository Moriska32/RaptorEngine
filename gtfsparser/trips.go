package gtfsparser

import (
	"EngineOT/raptor"
	"bufio"
	"io"
	"strconv"
	"strings"
)

//Stopsparser parsing gtfs stops
func TripsParser(pool io.Reader) {

	scanner := bufio.NewScanner(pool)
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Text()

		pool := strings.Split(line, ",")
		rid, _ := strconv.Atoi(pool[0])
		sid, _ := strconv.Atoi(pool[1])
		tid, _ := strconv.Atoi(pool[2])
		did, _ := strconv.Atoi(pool[4])
		raptor.Gtfs.Trips = append(raptor.Gtfs.Trips, raptor.Trip{
			Route_id:     rid,
			Service_id:   sid,
			Trip_id:      tid,
			Direction_id: did,
		})

	}

}
