package main
import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type WeatherResponse struct {
    Current struct {
        WeatherCode int `json:"weather_code"`
    } `json:"current"`
}

func weatherDescription(code int) string {
	switch code {
		case 0:
			return "Sunny"
		case 1, 2, 3:
			return "Cloudy"
		case 51, 53, 55:
			return "Drizzly"
		case 61, 63, 65:
			return "Rainy"
		case 80, 81, 82:
			return "Suddenly a rainy"
		case 95:
			return "Stormy"
		default:
			return " "
	}
}
func getWeather() string {
    response, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=14.5379338&longitude=121.045801&current=weather_code")

    if err != nil {
        fmt.Println("Something went wrong:", err)
        return "unknown"
    }

    defer response.Body.Close()

    fmt.Println(response.Status)
	body, err := io.ReadAll(response.Body)

	if err != nil {
        fmt.Println("Could not read response:", err)
        return "unknown"
	}
	var weather WeatherResponse

	err = json.Unmarshal(body, &weather)

	if err != nil {
    	fmt.Println("Could not parse JSON:", err)
    	return "unknown"
	}

	description := weatherDescription(weather.Current.WeatherCode)

	return description

}