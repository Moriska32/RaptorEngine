package gtfsparser

import (
	"bufio"
	"io"
)

//Stopsparser parsing gtfs stops
func RoutesParser(pool io.Reader) []Route {

	var routes []Route
	scanner := bufio.NewScanner(pool)
	scanner.Scan()
	for scanner.Scan() {

		line := scanner.Bytes()

		items := Re.FindAll(line, -1)
		pool := Replacer(items)
		routes = append(routes, Route{
			route_id:       pool[0],
			transport_type: pool[5],
		})

	}
	return routes
}
