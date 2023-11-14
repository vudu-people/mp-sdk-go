package model

type QR struct {
	Image            string `json:"image"`
	TemplateDocument string `json:"template_document"`
	TemplateImage    string `json:"template_image"`
}

type Pos struct {
	ID              int    `json:"id"`
	QR              QR     `json:"qr"`
	Status          string `json:"status"`
	DateCreated     string `json:"date_created"`
	DateLastUpdated string `json:"date_last_updated"`
	UUID            string `json:"uuid"`
	UserID          int    `json:"user_id"`
	Name            string `json:"name"`
	FixedAmount     bool   `json:"fixed_amount"`
	Category        int    `json:"category"`
	StoreID         string `json:"store_id"`
	ExternalID      string `json:"external_id"`
	Site            string `json:"site"`
	QrCode          string `json:"qr_code"`
}

type CreatePosRequest struct {
	StoreID         int    `json:"store_id"`
	Name            string `json:"name"`
	FixedAmount     bool   `json:"fixed_amount"`
	Category        int    `json:"category"`
	ExternalStoreID string `json:"external_store_id,omitempty"`
	ExternalID      string `json:"external_id,omitempty"`
}

type GetAllPosResponse struct {
	Paging struct {
		Total  int `json:"total"`
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"paging"`
	Results []Pos `json:"results"`
}
