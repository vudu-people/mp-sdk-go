package model

import "time"

type CreateStoreRequest struct {
	Name          string `json:"name"`
	BusinessHours struct {
		Monday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"monday,omitempty"`
		Tuesday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"tuesday,omitempty"`
		Wednesday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"wednesday,omitempty"`
		Thursday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"thursday,omitempty"`
		Friday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"friday,omitempty"`
		Saturday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"saturday,omitempty"`
		Sunday []struct {
			Open  string `json:"open"`
			Close string `json:"close"`
		} `json:"sunday,omitempty"`
	} `json:"business_hours"`
	Location struct {
		StreetNumber string  `json:"street_number"`
		StreetName   string  `json:"street_name"`
		CityName     string  `json:"city_name"`
		StateName    string  `json:"state_name"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
		Reference    string  `json:"reference"`
	} `json:"location"`
	ExternalID string `json:"external_id"`
}

type CreateStoreResponse struct {
	ID            int       `json:"id"`
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
	CreateStoreResponse
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
	Results []CreateStoreResponse `json:"results"`
}
