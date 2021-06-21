package gtfsparser

import (
	"EngineOT/raptor"
	"bufio"
	"io"
	"strconv"
)

//Stopsparser parsing gtfs stops
func RoutesParser(pool io.Reader) {

	scanner := bufio.NewScanner(pool)
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Bytes()

		items := raptor.Re.FindAll(line, -1)
		pool := Replacer(items)
		rid, _ := strconv.Atoi(pool[0])
		raptor.Routes = append(raptor.Routes, raptor.Route{
			Route_id:       rid,
			Transport_type: pool[5],
		})

	}

}
