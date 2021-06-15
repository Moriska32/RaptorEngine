package gtfsparser

import (
	"bufio"
	"io"
	"strings"
)

//Stopsparser parsing gtfs stops
func TripsParser(pool io.Reader) []Trips {

	var trips []Trips
	scanner := bufio.NewScanner(pool)
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Text()

		pool := strings.Split(line, ",")
		trips = append(trips, Trips{
			route_id:     pool[0],
			service_id:   pool[0],
			trip_id:      pool[0],
			direction_id: pool[0],
		})

	}
	return trips
}
