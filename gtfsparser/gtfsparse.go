package gtfsparser

import (
	"archive/zip"
	"fmt"
	"log"
	"time"
)

func GTFSParse() {

	//var stops []Stops

	// stopsf, err := ioutil.ReadFile("stops.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	Stops, StopTimes, Route, Trips := Unzip("gtfsparser\\gtfs\\gtfs.zip")
	_, _, _, _ = Stops, StopTimes, Route, Trips
	return
}
func Replacer(pool [][]byte) []string {

	res := []string{}
	for _, item := range pool {
		item := string(item)

		if string(item[0]) == "," {
			item = item[1:]
		}
		res = append(res, item)

	}
	return res

}
func Unzip(path string) ([]Stops, []StopTimes, []Route, []Trips) {
	start := time.Now()
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	var stops []Stops
	var stoptimes []StopTimes
	var routes []Route
	var trips []Trips
	for _, f := range r.File {
		switch {
		case f.Name == "stops.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			stops = StopsParser(rc)
			fmt.Println(len(stops))

		case f.Name == "stop_times.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			stoptimes = StopTimesParser(rc)
			fmt.Println(len(stoptimes))
		case f.Name == "routes.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			routes = RoutesParser(rc)
			fmt.Println(len(routes))

		case f.Name == "trips.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			trips = TripsParser(rc)
			fmt.Println(len(trips))
		}

	}
	duration := time.Since(start)
	fmt.Println(duration)
	return stops, stoptimes, routes, trips
}
