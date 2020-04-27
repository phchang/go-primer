package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sketch-go-course/pkg/location"
	"time"
)

type Forecast struct {
	Properties struct {
		Periods []Period
	}
}

func (f Forecast) Summary() ForecastSummary {

	daysMap := make(map[string]ForecastDay)

	for _, p := range f.Properties.Periods {

		weekday := p.StartTime.Weekday().String()

		day, ok := daysMap[weekday]
		if !ok {
			daysMap[weekday] = ForecastDay{
				Day:           p.StartTime,
				Low:           p.Temperature,
				High:          p.Temperature,
				ShortForecast: p.ShortForecast,
			}
			continue
		}

		if p.Temperature > day.Low {
			day.High = p.Temperature
		} else {
			day.Low = p.Temperature
		}

		day.ShortForecast += "; " + p.ShortForecast

		daysMap[weekday] = day
	}

	days := make([]ForecastDay, 0, len(daysMap))

	for _, day := range daysMap {
		days = append(days, day)
	}

	return ForecastSummary{
		Days: days,
	}

}

type ForecastSummary struct {
	Days []ForecastDay
}

type ForecastDay struct {
	Day           time.Time
	Low           float64
	High          float64
	ShortForecast string
}

type Points struct {
	Properties struct {
		ForecastURL string `json:"forecast"`
	}
}

type Period struct {
	Number        float64
	Name          string
	StartTime     time.Time
	Temperature   float64
	ShortForecast string
}

func (p *Period) UnmarshalJSON(data []byte) error {

	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	a, _ := v.(map[string]interface{})

	p.Number, _ = a["number"].(float64)
	p.Name, _ = a["name"].(string)
	p.Temperature, _ = a["temperature"].(float64)
	p.ShortForecast, _ = a["shortForecast"].(string)

	startTimeStr, _ := a["startTime"].(string)

	// "2020-04-24T18:00:00-04:00"
	t, _ := time.Parse("2006-01-02T15:04:05-07:00", startTimeStr)
	p.StartTime = t

	return nil
}

type Client struct {
	Client *http.Client
}

func (c Client) FetchForecast(coordinates location.Coordinate) (Forecast, error) {

	httpClient := c.Client

	if res, getErr := httpClient.Get("https://api.weather.gov/points/" + coordinates.String()); getErr != nil {
		fmt.Println("error calling GET", getErr)
		return Forecast{}, getErr

	} else {
		fmt.Println("response = ", res.StatusCode)

		bodyBytes, _ := ioutil.ReadAll(res.Body)

		var points Points
		_ = json.Unmarshal(bodyBytes, &points)

		if res, getErr := httpClient.Get(points.Properties.ForecastURL); getErr != nil {
			fmt.Println("error calling GET", getErr)
			return Forecast{}, getErr

		} else {
			bodyBytes, _ := ioutil.ReadAll(res.Body)
			bodyString := string(bodyBytes)

			fmt.Println(bodyString)

			var forecast Forecast
			_ = json.Unmarshal(bodyBytes, &forecast)
			return forecast, nil
		}
	}
}
