package gtfsparser

import (
	"EngineOT/raptor"
	"bufio"
	"io"
	"strconv"
)

//Stopsparser parsing gtfs stops
func StopsParser(pool io.Reader) {

	scanner := bufio.NewScanner(pool)
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Bytes()

		items := raptor.Re.FindAll(line, -1)
		pool := Replacer(items)
		sid, _ := strconv.Atoi(pool[0])
		sc, _ := strconv.Atoi(pool[1])
		lat, _ := strconv.ParseFloat(pool[3], 64)
		lon, _ := strconv.ParseFloat(pool[4], 64)
		lt, _ := strconv.Atoi(pool[5])
		wb, _ := strconv.Atoi(pool[6])
		raptor.Stops = append(raptor.Stops, raptor.Stop{
			Stop_id:             sid,
			Stop_code:           sc,
			Stop_name:           pool[2],
			Stop_lat:            lat,
			Stop_lon:            lon,
			Location_type:       lt,
			Wheelchair_boarding: wb,
			Transport_type:      pool[7],
		})

	}

}
