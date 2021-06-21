package gtfsparser

import (
	"EngineOT/raptor"
	"bufio"
	"fmt"
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
		sid, err := strconv.Atoi(pool[0])
		if err != nil {
			fmt.Sprintln(err)
		}
		tid, err := strconv.Atoi(pool[1])
		if err != nil {
			fmt.Sprintln(err)
		}
		raptor.Gtfs.StopTimes = append(raptor.Gtfs.StopTimes, raptor.StopTime{
			Stop_id:        sid,
			Trip_id:        tid,
			Arrival_time:   pool[2],
			Departure_time: pool[3],
		})

	}

}
