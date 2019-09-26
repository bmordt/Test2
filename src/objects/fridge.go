package objects

type Fridge struct {
	Id          string  `json:"id"`
	Timestamp   int     `json:"timestamp"`
	Temperature float64 `json:"temperature"`
}
type FridgeList struct {
	FridgeList []Fridge
}

type FridgeResult struct {
	Id      string    `json:"id"`
	Average float64   `json:"average"`
	Median  float64   `json:"Median"`
	Mode    []float64 `json:"Mode"`
}

func GetTemperatureListForFridge(fridges []Fridge) []float64 {
	var temps []float64
	for _, fridge := range fridges {
		temps = append(temps, fridge.Temperature)
	}
	return temps
}
