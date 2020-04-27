package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sketch-go-course/pkg/location"
	"sketch-go-course/pkg/weather"
)

func main() {

	router := mux.NewRouter()

	zipCodeMap, _ := location.LoadZipCodeMap("zip.csv")

	weatherClient := weather.Client{
		Client: &http.Client{},
	}

	router.HandleFunc("/forecast/{zipcode}", func(writer http.ResponseWriter, request *http.Request) {

		vars := mux.Vars(request)
		zip := vars["zipcode"]
		coords := zipCodeMap[zip]

		forecast, fetchErr := weatherClient.FetchForecast(coords)

		if fetchErr != nil {
			fmt.Println("Could not get forecast ", fetchErr)
			return
		}

		b, _ := json.Marshal(forecast.Summary())

		writer.Header().Add("content-type", "application/json")
		_, _ = writer.Write(b)

	})

	server := &http.Server{Handler: router,
		Addr: ":8000"}
	server.ListenAndServe()

}
