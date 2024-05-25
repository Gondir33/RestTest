package models

type Currency struct {
	ID           int     `json:"id"`
	NameCurrency string  `json:"model"`
	Value        float64 `json:"value"`
}
