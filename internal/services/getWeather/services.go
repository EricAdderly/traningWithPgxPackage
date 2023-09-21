package getWeatherServices

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/traningWithPgxPackage/internal/models"
)

func GetMeWeather(town string, client *redis.Client) (*models.Weather, error) {

	value, err := client.Get(town).Bytes()
	if err != nil {
		currentWeather, err := makeRequest(town, client)
		if err != nil {
			return nil, fmt.Errorf("GetWeather: %w", err)
		}

		return currentWeather, nil
	}

	var currentWeather *models.Weather
	err = json.Unmarshal(value, &currentWeather)
	if err != nil {
		return nil, fmt.Errorf("GetWeather: %w", err)
	}

	return currentWeather, nil
}

func makeRequest(town string, client *redis.Client) (*models.Weather, error) {
	fmt.Println(123)
	tempURL := newLink(town)
	resp, err := http.Get(tempURL)
	if err != nil {
		return nil, fmt.Errorf("GetWeather: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GetWeather: %w", err)
	}

	var currentWeather *models.Weather
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		return nil, fmt.Errorf("GetWeather: %w", err)
	}

	data, err := json.Marshal(currentWeather)
	if err != nil {
		return nil, fmt.Errorf("GetWeather: %w", err)
	}

	err = client.Set(town, data, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("GetWeather: %w", err)
	}

	return currentWeather, nil
}
func newLink(town string) string {
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=697381875f993add6f43195450798d80&units=metric", town)
}
