package dto

import "time"

type Transaction struct {
	ID              string
	Name            string
	Number          string
	ExperationMonth int32
	ExperationYear  int32
	CVV             int32
	Amount          float64
	Store           string
	Description     string
	CreatedAt       time.Time
}