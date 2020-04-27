package weather

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"sketch-go-course/pkg/location"
	"testing"
)

var mockResponse = `
{
    "@context": [
        "https://raw.githubusercontent.com/geojson/geojson-ld/master/contexts/geojson-base.jsonld",
        {
            "wx": "https://api.weather.gov/ontology#",
            "s": "https://schema.org/",
            "geo": "http://www.opengis.net/ont/geosparql#",
            "unit": "http://codes.wmo.int/common/unit/",
            "@vocab": "https://api.weather.gov/ontology#",
            "geometry": {
                "@id": "s:GeoCoordinates",
                "@type": "geo:wktLiteral"
            },
            "city": "s:addressLocality",
            "state": "s:addressRegion",
            "distance": {
                "@id": "s:Distance",
                "@type": "s:QuantitativeValue"
            },
            "bearing": {
                "@type": "s:QuantitativeValue"
            },
            "value": {
                "@id": "s:value"
            },
            "unitCode": {
                "@id": "s:unitCode",
                "@type": "@id"
            },
            "forecastOffice": {
                "@type": "@id"
            },
            "forecastGridData": {
                "@type": "@id"
            },
            "publicZone": {
                "@type": "@id"
            },
            "county": {
                "@type": "@id"
            }
        }
    ],
    "id": "https://api.weather.gov/points/18.1805999,-66.75",
    "type": "Feature",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -66.75,
            18.180599900000001
        ]
    },
    "properties": {
        "@id": "https://api.weather.gov/points/18.1805999,-66.75",
        "@type": "wx:Point",
        "cwa": "SJU",
        "forecastOffice": "https://api.weather.gov/offices/SJU",
        "gridX": 107,
        "gridY": 106,
        "forecast": "https://api.weather.gov/gridpoints/SJU/107,106/forecast",
        "forecastHourly": "https://api.weather.gov/gridpoints/SJU/107,106/forecast/hourly",
        "forecastGridData": "https://api.weather.gov/gridpoints/SJU/107,106",
        "observationStations": "https://api.weather.gov/gridpoints/SJU/107,106/stations",
        "relativeLocation": {
            "type": "Feature",
            "geometry": {
                "type": "Point",
                "coordinates": [
                    -66.723544000000004,
                    18.163775999999999
                ]
            },
            "properties": {
                "city": "Adjuntas",
                "state": "PR",
                "distance": {
                    "value": 3363.3262476376262,
                    "unitCode": "unit:m"
                },
                "bearing": {
                    "value": 303,
                    "unitCode": "unit:degrees_true"
                }
            }
        },
        "forecastZone": "https://api.weather.gov/zones/forecast/PRZ009",
        "county": "https://api.weather.gov/zones/county/PRC001",
        "fireWeatherZone": "https://api.weather.gov/zones/fire/PRZ023",
        "timeZone": "America/Puerto_Rico",
        "radarStation": "TJUA"
    }
}
`

var mockResponse2 = `
{
    "@context": [
        "https://raw.githubusercontent.com/geojson/geojson-ld/master/contexts/geojson-base.jsonld",
        {
            "wx": "https://api.weather.gov/ontology#",
            "geo": "http://www.opengis.net/ont/geosparql#",
            "unit": "http://codes.wmo.int/common/unit/",
            "@vocab": "https://api.weather.gov/ontology#"
        }
    ],
    "type": "Feature",
    "geometry": {
        "type": "GeometryCollection",
        "geometries": [
            {
                "type": "Point",
                "coordinates": [
                    -66.747802699999994,
                    18.186239199999999
                ]
            },
            {
                "type": "Polygon",
                "coordinates": [
                    [
                        [
                            -66.753787900000006,
                            18.191919299999999
                        ],
                        [
                            -66.753787900000006,
                            18.1805591
                        ],
                        [
                            -66.741817400000002,
                            18.1805591
                        ],
                        [
                            -66.741817400000002,
                            18.191919299999999
                        ],
                        [
                            -66.753787900000006,
                            18.191919299999999
                        ]
                    ]
                ]
            }
        ]
    },
    "properties": {
        "updated": "2020-04-24T13:40:02+00:00",
        "units": "us",
        "forecastGenerator": "BaselineForecastGenerator",
        "generatedAt": "2020-04-24T15:56:04+00:00",
        "updateTime": "2020-04-24T13:40:02+00:00",
        "validTimes": "2020-04-24T07:00:00+00:00/P8DT6H",
        "elevation": {
            "value": 467.86800000000005,
            "unitCode": "unit:m"
        },
        "periods": [
            {
                "number": 1,
                "name": "Today",
                "startTime": "2020-04-24T11:00:00-04:00",
                "endTime": "2020-04-24T18:00:00-04:00",
                "isDaytime": true,
                "temperature": 87,
                "temperatureUnit": "F",
                "temperatureTrend": "falling",
                "windSpeed": "10 to 14 mph",
                "windDirection": "ESE",
                "icon": "https://api.weather.gov/icons/land/day/sct/rain_showers,30?size=medium",
                "shortForecast": "Mostly Sunny then Scattered Rain Showers",
                "detailedForecast": "Scattered rain showers after noon. Mostly sunny. High near 87, with temperatures falling to around 82 in the afternoon. East southeast wind 10 to 14 mph. Chance of precipitation is 30%."
            },
            {
                "number": 2,
                "name": "Tonight",
                "startTime": "2020-04-24T18:00:00-04:00",
                "endTime": "2020-04-25T06:00:00-04:00",
                "isDaytime": false,
                "temperature": 70,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "7 to 10 mph",
                "windDirection": "ESE",
                "icon": "https://api.weather.gov/icons/land/night/rain_showers/few?size=medium",
                "shortForecast": "Isolated Rain Showers then Mostly Clear",
                "detailedForecast": "Isolated rain showers before 9pm. Mostly clear, with a low around 70. East southeast wind 7 to 10 mph."
            },
            {
                "number": 3,
                "name": "Saturday",
                "startTime": "2020-04-25T06:00:00-04:00",
                "endTime": "2020-04-25T18:00:00-04:00",
                "isDaytime": true,
                "temperature": 88,
                "temperatureUnit": "F",
                "temperatureTrend": "falling",
                "windSpeed": "10 to 14 mph",
                "windDirection": "SE",
                "icon": "https://api.weather.gov/icons/land/day/few/rain_showers,20?size=medium",
                "shortForecast": "Sunny then Isolated Rain Showers",
                "detailedForecast": "Isolated rain showers after noon. Sunny. High near 88, with temperatures falling to around 83 in the afternoon. Southeast wind 10 to 14 mph. Chance of precipitation is 20%."
            },
            {
                "number": 4,
                "name": "Saturday Night",
                "startTime": "2020-04-25T18:00:00-04:00",
                "endTime": "2020-04-26T06:00:00-04:00",
                "isDaytime": false,
                "temperature": 70,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "6 to 12 mph",
                "windDirection": "E",
                "icon": "https://api.weather.gov/icons/land/night/few?size=medium",
                "shortForecast": "Mostly Clear",
                "detailedForecast": "Mostly clear, with a low around 70. East wind 6 to 12 mph."
            },
            {
                "number": 5,
                "name": "Sunday",
                "startTime": "2020-04-26T06:00:00-04:00",
                "endTime": "2020-04-26T18:00:00-04:00",
                "isDaytime": true,
                "temperature": 88,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "12 to 16 mph",
                "windDirection": "E",
                "icon": "https://api.weather.gov/icons/land/day/sct/rain_showers,50?size=medium",
                "shortForecast": "Mostly Sunny then Scattered Rain Showers",
                "detailedForecast": "Scattered rain showers after noon. Mostly sunny, with a high near 88. East wind 12 to 16 mph. Chance of precipitation is 50%."
            },
            {
                "number": 6,
                "name": "Sunday Night",
                "startTime": "2020-04-26T18:00:00-04:00",
                "endTime": "2020-04-27T06:00:00-04:00",
                "isDaytime": false,
                "temperature": 68,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "7 to 10 mph",
                "windDirection": "E",
                "icon": "https://api.weather.gov/icons/land/night/skc?size=medium",
                "shortForecast": "Clear",
                "detailedForecast": "Clear, with a low around 68. East wind 7 to 10 mph."
            },
            {
                "number": 7,
                "name": "Monday",
                "startTime": "2020-04-27T06:00:00-04:00",
                "endTime": "2020-04-27T18:00:00-04:00",
                "isDaytime": true,
                "temperature": 86,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "7 to 12 mph",
                "windDirection": "ESE",
                "icon": "https://api.weather.gov/icons/land/day/rain_showers/rain_showers,30?size=medium",
                "shortForecast": "Scattered Rain Showers",
                "detailedForecast": "Scattered rain showers. Mostly sunny, with a high near 86. East southeast wind 7 to 12 mph. Chance of precipitation is 30%."
            },
            {
                "number": 8,
                "name": "Monday Night",
                "startTime": "2020-04-27T18:00:00-04:00",
                "endTime": "2020-04-28T06:00:00-04:00",
                "isDaytime": false,
                "temperature": 68,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "7 mph",
                "windDirection": "ESE",
                "icon": "https://api.weather.gov/icons/land/night/rain_showers?size=medium",
                "shortForecast": "Isolated Rain Showers",
                "detailedForecast": "Isolated rain showers. Mostly clear, with a low around 68. East southeast wind around 7 mph."
            },
            {
                "number": 9,
                "name": "Tuesday",
                "startTime": "2020-04-28T06:00:00-04:00",
                "endTime": "2020-04-28T18:00:00-04:00",
                "isDaytime": true,
                "temperature": 86,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "8 mph",
                "windDirection": "E",
                "icon": "https://api.weather.gov/icons/land/day/rain_showers/rain_showers,30?size=medium",
                "shortForecast": "Scattered Rain Showers",
                "detailedForecast": "Scattered rain showers. Mostly sunny, with a high near 86. East wind around 8 mph. Chance of precipitation is 30%."
            },
            {
                "number": 10,
                "name": "Tuesday Night",
                "startTime": "2020-04-28T18:00:00-04:00",
                "endTime": "2020-04-29T06:00:00-04:00",
                "isDaytime": false,
                "temperature": 68,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "7 mph",
                "windDirection": "E",
                "icon": "https://api.weather.gov/icons/land/night/rain_showers?size=medium",
                "shortForecast": "Isolated Rain Showers",
                "detailedForecast": "Isolated rain showers. Mostly clear, with a low around 68."
            },
            {
                "number": 11,
                "name": "Wednesday",
                "startTime": "2020-04-29T06:00:00-04:00",
                "endTime": "2020-04-29T18:00:00-04:00",
                "isDaytime": true,
                "temperature": 86,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "9 mph",
                "windDirection": "E",
                "icon": "https://api.weather.gov/icons/land/day/rain_showers/rain_showers,50?size=medium",
                "shortForecast": "Scattered Rain Showers",
                "detailedForecast": "Scattered rain showers. Mostly sunny, with a high near 86. Chance of precipitation is 50%."
            },
            {
                "number": 12,
                "name": "Wednesday Night",
                "startTime": "2020-04-29T18:00:00-04:00",
                "endTime": "2020-04-30T06:00:00-04:00",
                "isDaytime": false,
                "temperature": 68,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "7 mph",
                "windDirection": "E",
                "icon": "https://api.weather.gov/icons/land/night/rain_showers?size=medium",
                "shortForecast": "Isolated Rain Showers",
                "detailedForecast": "Isolated rain showers. Mostly clear, with a low around 68."
            },
            {
                "number": 13,
                "name": "Thursday",
                "startTime": "2020-04-30T06:00:00-04:00",
                "endTime": "2020-04-30T18:00:00-04:00",
                "isDaytime": true,
                "temperature": 86,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "8 to 12 mph",
                "windDirection": "ESE",
                "icon": "https://api.weather.gov/icons/land/day/rain_showers/rain_showers,40?size=medium",
                "shortForecast": "Scattered Rain Showers",
                "detailedForecast": "Scattered rain showers. Mostly sunny, with a high near 86. Chance of precipitation is 40%."
            },
            {
                "number": 14,
                "name": "Thursday Night",
                "startTime": "2020-04-30T18:00:00-04:00",
                "endTime": "2020-05-01T06:00:00-04:00",
                "isDaytime": false,
                "temperature": 69,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "7 to 10 mph",
                "windDirection": "ESE",
                "icon": "https://api.weather.gov/icons/land/night/rain_showers?size=medium",
                "shortForecast": "Isolated Rain Showers",
                "detailedForecast": "Isolated rain showers. Mostly clear, with a low around 69."
            }
        ]
    }
}`

type MockClient struct {
	Fn func(*http.Request) (*http.Response, error)
}

func (c MockClient) RoundTrip(req *http.Request) (*http.Response, error) {
	return c.Fn(req)
}

func TestFetchForecast(t *testing.T) {

	c := Client{
		Client: &http.Client{
			Transport: MockClient{
				Fn: func(request *http.Request) (*http.Response, error) {

					switch {
					case request.URL.String() == "https://api.weather.gov/points/38.676026,-90.377994":
						res := &http.Response{}
						res.StatusCode = 200
						res.Body = ioutil.NopCloser(bytes.NewReader([]byte(mockResponse)))
						return res, nil
					case request.URL.String() == "https://api.weather.gov/gridpoints/SJU/107,106/forecast":
						res := &http.Response{}
						res.StatusCode = 200
						res.Body = ioutil.NopCloser(bytes.NewReader([]byte(mockResponse2)))
						return res, nil
					}
					return nil, nil
				},
			},
		},
	}

	forecast, err := c.FetchForecast(location.Coordinate{
		Lat:  "38.676026",
		Long: "-90.377994",
	})

	require.NoError(t, err)
	assert.Len(t, forecast.Properties.Periods, 14)
}
