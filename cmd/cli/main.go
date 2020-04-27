package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sketch-go-course/pkg/location"
	"sketch-go-course/pkg/weather"
	"sort"
	"strings"
)

func main() {

	// 1. type in zip code at the command prompt
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter zip code: ")
	zipCodeStr, _ := reader.ReadString('\n')

	// 2. get the forecast using the entered zip code

	//  - parse CSV

	zipCodeMap, zipCodeErr := location.LoadZipCodeMap("zip.csv")

	if zipCodeErr != nil {
		fmt.Println("Could not load zip.csv", zipCodeStr)
		return
	}

	coords, ok := zipCodeMap[strings.TrimSpace(zipCodeStr)]

	if !ok {
		fmt.Println("Could not find zipcode ", zipCodeStr)
		return
	}

	weatherClient := weather.Client{
		Client: &http.Client{},
	}

	forecast, fetchErr := weatherClient.FetchForecast(coords)

	if fetchErr != nil {
		fmt.Println("Could not get forecast ", fetchErr)
		return
	}

	fmt.Println("\nForecast for ", zipCodeStr)

	summary := forecast.Summary()

	sort.Slice(summary.Days, func(i, j int) bool {
		return summary.Days[i].Day.Before(summary.Days[j].Day)
	})

	for _, day := range summary.Days {
		fmt.Printf("%v\t\t%v\t%v\t%v\n", day.Day.Weekday().String(), day.Low, day.High, day.ShortForecast)
	}
}
