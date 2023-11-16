package model

type GetPaymentMethodsResponse []struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PaymentTypeID   string `json:"payment_type_id"`
	Status          string `json:"status"`
	SecureThumbnail string `json:"secure_thumbnail"`
	Thumbnail       string `json:"thumbnail"`
	DeferredCapture string `json:"deferred_capture"`
	Settings        struct {
		Bin struct {
			Pattern             string `json:"pattern"`
			ExclusionPattern    string `json:"exclusion_pattern"`
			InstallmentsPattern string `json:"installments_pattern"`
		} `json:"bin"`
		CardNumber struct {
			Length     int    `json:"length"`
			Validation string `json:"validation"`
		} `json:"card_number"`
		SecurityCode struct {
			Mode         string `json:"mode"`
			Length       int    `json:"length"`
			CardLocation string `json:"card_location"`
		} `json:"security_code"`
	} `json:"settings"`
	AdditionalInfoNeeded []struct {
	} `json:"additional_info_needed"`
	MinAllowedAmount      float64 `json:"min_allowed_amount"`
	MaxAllowedAmount      int     `json:"max_allowed_amount"`
	AccreditationTime     int     `json:"accreditation_time"`
	FinancialInstitutions struct {
	} `json:"financial_institutions"`
	ProcessingModes string `json:"processing_modes"`
}
