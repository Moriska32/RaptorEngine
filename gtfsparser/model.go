package gtfsparser

import "regexp"

var Re = regexp.MustCompile(`(?:,|\n|^)("(?:(?:"")*[^"]*)*"|[^",\n]*|(?:\n|$))`)

//StopTime stoptime
type StopTimes struct {
	stop_id        string
	trip_id        string
	arrival_time   string
	departure_time string
}

//Trips trips
type Trips struct {
	route_id     string
	service_id   string
	trip_id      string
	direction_id string
}

//Route routes
type Route struct {
	route_id       string
	transport_type string
}

//Stops stops
type Stops struct {
	stop_id             string
	stop_code           string
	stop_name           string
	stop_lat            string
	stop_lon            string
	location_type       string
	wheelchair_boarding string
	transport_type      string
}
