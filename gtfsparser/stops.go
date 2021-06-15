package gtfsparser

import (
	"bufio"
	"io"
)

//Stopsparser parsing gtfs stops
func StopsParser(pool io.Reader) []Stops {

	var stops []Stops
	scanner := bufio.NewScanner(pool)
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Bytes()

		items := Re.FindAll(line, -1)
		pool := Replacer(items)
		stops = append(stops, Stops{
			stop_id:             pool[0],
			stop_code:           pool[1],
			stop_name:           pool[2],
			stop_lat:            pool[3],
			stop_lon:            pool[4],
			location_type:       pool[5],
			wheelchair_boarding: pool[6],
			transport_type:      pool[7],
		})

	}
	return stops
}
