package main

import (
	"EngineOT/gtfsparser"
	"EngineOT/raptor"
)

func main() {

	gtfsparser.GTFSParse()

	from_stop := "19306"
	to_stop := "22481"
	_, _ = from_stop, to_stop
	raptor.GetStops(from_stop, to_stop)
}
