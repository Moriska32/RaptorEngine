package main

import (
	"EngineOT/gtfsparser"
	"EngineOT/raptor"
	"fmt"
)

func main() {

	if err := gtfsparser.Unzip("gtfsparser\\gtfs\\gtfs.zip"); err != nil {
		fmt.Sprintln(err)
	}

	from_stop := "19306"
	to_stop := "22481"

	raptor.GetStops(from_stop, to_stop)
	fmt.Println(raptor.C)
}
