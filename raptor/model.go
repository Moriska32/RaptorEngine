package raptor

import "regexp"

var (
	Re = regexp.MustCompile(`(?:,|\n|^)("(?:(?:"")*[^"]*)*"|[^",\n]*|(?:\n|$))`)

	C    CoordsOfStops
	Gtfs GtfsStruct
)

type GtfsStruct struct {
	Stops     []Stop
	StopTimes []StopTime
	Trips     []Trip
	Routes    []Route
}
type CoordsOfStops struct {
	FromStopLon float64
	FromStopLat float64
	ToStopLon   float64
	ToStopLat   float64
}

//StopTime stoptime
type StopTime struct {
	Stop_id        int
	Trip_id        int
	Arrival_time   string
	Departure_time string
}

//Trips trips
type Trip struct {
	Route_id     int
	Service_id   int
	Trip_id      int
	Direction_id int
}

//Route routes
type Route struct {
	Route_id       int
	Transport_type string
}

//Stop stops
type Stop struct {
	Stop_id             int
	Stop_code           int
	Stop_name           string
	Stop_lat            float64
	Stop_lon            float64
	Location_type       int
	Wheelchair_boarding int
	Transport_type      string
}
