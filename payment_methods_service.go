package mpsdkgo

import (
	"fmt"
	"net/http"

	"github.com/vudu-people/mp-sdk-go/model"
)

type PaymentMethodsService struct {
	client *ApiClient
}

type IPaymentMethodsService interface {
	GetPaymentMethods() (model.GetPaymentMethodsResponse, error)
}

func NewPaymentMethodsService(client *ApiClient) *PaymentMethodsService {
	return &PaymentMethodsService{client}
}

func (s *PaymentMethodsService) GetPaymentMethods() (model.GetPaymentMethodsResponse, error) {

	var (
		path = "%s/v1/payment_methods"
	)

	url := fmt.Sprintf(path, s.client.BaseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.GetPaymentMethodsResponse{}, err
	}

	r, err := s.client.callAPI(req)
	if err != nil {
		return model.GetPaymentMethodsResponse{}, err
	}

	defer r.Body.Close()

	var response model.GetPaymentMethodsResponse
	s.client.DeserializeBody(r.Body, &response)

	return response, nil

}
