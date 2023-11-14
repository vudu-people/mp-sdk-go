package model

import "time"

type BusinessHourPair struct {
	Open  string `json:"open"`
	Close string `json:"close"`
}

type BusinessHours struct {
	Monday    []BusinessHourPair `json:"monday,omitempty"`
	Tuesday   []BusinessHourPair `json:"tuesday,omitempty"`
	Wednesday []BusinessHourPair `json:"wednesday,omitempty"`
	Thursday  []BusinessHourPair `json:"thursday,omitempty"`
	Friday    []BusinessHourPair `json:"friday,omitempty"`
	Saturday  []BusinessHourPair `json:"saturday,omitempty"`
	Sunday    []BusinessHourPair `json:"sunday,omitempty"`
}

type CreateStoreRequest struct {
	Name          string        `json:"name,omitempty"`
	BusinessHours BusinessHours `json:"business_hours"`
	Location      struct {
		StreetNumber string  `json:"street_number,omitempty"`
		StreetName   string  `json:"street_name,omitempty"`
		CityName     string  `json:"city_name,omitempty"`
		StateName    string  `json:"state_name,omitempty"`
		Latitude     float64 `json:"latitude,omitempty"`
		Longitude    float64 `json:"longitude,omitempty"`
		Reference    string  `json:"reference,omitempty"`
	} `json:"location,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

type Store struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	DateCreated   time.Time `json:"date_created"`
	BusinessHours struct {
		Monday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"monday"`
		Tuesday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"tuesday"`
	} `json:"business_hours"`
	Location struct {
		AddressLine string  `json:"address_line"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		Reference   string  `json:"reference"`
	} `json:"location"`
	ExternalID string `json:"external_id"`
}

type GetStoreResponse struct {
	Store
}

type UpdateStoreRequest struct {
	CreateStoreRequest
}

type GetStoresResponse struct {
	Paging struct {
		Total  int `json:"total"`
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"paging"`
	Results []Store `json:"results"`
}
