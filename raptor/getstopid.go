package raptor

import (
	"fmt"
	"strconv"
)

func GetStops(from string, to string) {

	from_int, err1 := strconv.Atoi(from)

	to_int, err2 := strconv.Atoi(to)

	switch {
	case err1 == nil && err2 == nil:

		C.GetStopCoordsByID(from_int, to_int)
	case err1 != nil && err2 != nil:
		C.GetStopCoordsByName(from, to)

	case (err1 == nil && err2 != nil) || (err1 != nil && err2 == nil):

		fmt.Println("Wrong format")
	}

}
func (C *CoordsOfStops) GetStopCoordsByID(from int, to int) {

	for _, item := range Gtfs.Stops {

		if item.Stop_id == from {
			C.FromStopLat = item.Stop_lat
			C.FromStopLon = item.Stop_lon
		} else if item.Stop_id == to {

			C.ToStopLat = item.Stop_lat
			C.ToStopLon = item.Stop_lon

		}

	}

}
func (C *CoordsOfStops) GetStopCoordsByName(from string, to string) {

	for _, item := range Gtfs.Stops {

		if item.Stop_name == from {
			C.FromStopLat = item.Stop_lat
			C.FromStopLon = item.Stop_lon
		} else if item.Stop_name == to {

			C.ToStopLat = item.Stop_lat
			C.ToStopLon = item.Stop_lon

		}

	}

}
