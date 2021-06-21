package gtfsparser

import (
	"EngineOT/raptor"
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
	Unzip("gtfsparser\\gtfs\\gtfs.zip")

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
func Unzip(path string) {
	start := time.Now()
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		switch {
		case f.Name == "stops.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			StopsParser(rc)
			fmt.Println(len(raptor.Stops))

		case f.Name == "stop_times.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			StopTimesParser(rc)
			fmt.Println(len(raptor.StopTimes))
		case f.Name == "routes.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			RoutesParser(rc)
			fmt.Println(len(raptor.Routes))

		case f.Name == "trips.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			TripsParser(rc)
			fmt.Println(len(raptor.Trips))
		}

	}
	duration := time.Since(start)
	fmt.Println(duration)

}
