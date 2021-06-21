package gtfsparser

import (
	"EngineOT/raptor"
	"bufio"
	"io"
	"strconv"
	"strings"
)

//Stopsparser parsing gtfs stops
func StopTimesParser(pool io.Reader) {

	scanner := bufio.NewScanner(pool)
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Text()

		pool := strings.Split(line, ",")
		sid, _ := strconv.Atoi(pool[0])
		tid, _ := strconv.Atoi(pool[1])
		raptor.StopTimes = append(raptor.StopTimes, raptor.StopTime{
			Stop_id:        sid,
			Trip_id:        tid,
			Arrival_time:   pool[2],
			Departure_time: pool[3],
		})

	}

}
