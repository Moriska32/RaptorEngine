package gtfsparser

import (
	"bufio"
	"io"
	"strings"
)

//Stopsparser parsing gtfs stops
func StopTimesParser(pool io.Reader) []StopTimes {

	var stops []StopTimes
	scanner := bufio.NewScanner(pool)
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Text()

		pool := strings.Split(line, ",")
		stops = append(stops, StopTimes{
			stop_id:        pool[0],
			trip_id:        pool[1],
			arrival_time:   pool[2],
			departure_time: pool[3],
		})

	}
	return stops
}
