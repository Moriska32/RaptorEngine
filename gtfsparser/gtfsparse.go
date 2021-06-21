package gtfsparser

import (
	"EngineOT/raptor"
	"archive/zip"
	"fmt"
	"log"
)

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
func Unzip(path string) error {

	//start := time.Now()
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
				fmt.Println(err)
			}

			StopsParser(rc)
			fmt.Println(len(raptor.Gtfs.Stops))

		case f.Name == "stop_times.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				fmt.Println(err)
			}
			StopTimesParser(rc)
			fmt.Println(len(raptor.Gtfs.StopTimes))

		case f.Name == "routes.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				fmt.Println(err)
			}

			RoutesParser(rc)
			fmt.Println(len(raptor.Gtfs.Routes))

		case f.Name == "trips.txt":
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				fmt.Println(err)
			}

			TripsParser(rc)
			fmt.Println(len(raptor.Gtfs.Trips))
		}

	}
	//duration := time.Since(start)
	//fmt.Println(duration)

	return err
}
