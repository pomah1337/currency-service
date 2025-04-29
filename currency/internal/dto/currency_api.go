package dto

type ResponseData struct {
	Date string             `json:"date"`
	Rub  map[string]float64 `json:"rub"`
}
