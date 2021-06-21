package raptor

import "strconv"

func GetStops(from string, to string) {

	from_int, err1 := strconv.Atoi("from")
	to_int, err2 := strconv.Atoi("to")

	if err1 == nil && err2 == nil {
		_, _ = from_int, to_int
	}

}
func GetStopID(stops *Stop, from int, to int) {

}
