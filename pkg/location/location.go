package location

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Coordinate struct {
	Lat  string
	Long string
}

func (c Coordinate) String() string {
	return c.Lat + "," + c.Long
}

func LoadZipCodeMap(filename string) (zipCodeMap map[string]Coordinate, err error) {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Could not open zip.csv", err)
		return nil, err
	}

	csvReader := csv.NewReader(f)

	records, csvReadErr := csvReader.ReadAll()

	if csvReadErr != nil {
		fmt.Println("Could not read zip.csv")
		return
	}

	zipCodeMap = make(map[string]Coordinate)

	for _, record := range records[1:] {

		zipCodeMap[record[0]] = Coordinate{
			Lat:  record[1],
			Long: record[2],
		}
	}

	return
}
