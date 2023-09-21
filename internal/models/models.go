package models

// тут лежат структурки
type Weather struct {
	Main MainJson `json:"main"`
	Wind WindJson `json:"wind"`
}

type MainJson struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
}
type WindJson struct {
	WindSpeed float32 `json:"speed"`
}
