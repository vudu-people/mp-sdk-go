package mpsdkgo

//go:generate mockgen -source=store_service.go -destination=mocks/store_service_mock.go -package=mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/vudu-people/mp-sdk-go/model"
)

type StoreService struct {
	client *ApiClient
}

type IStoreService interface {
	CreateStore(userID string, body model.CreateStoreRequest) (model.Store, error)
	GetStore(storeID string) (model.GetStoreResponse, error)
	GetStoreByExternalID(userID string, externalID string) (model.GetStoresResponse, error)
	GetStores(userID string, limit int) (model.GetStoresResponse, error)
	UpdateStore(userID string, storeID string, body model.UpdateStoreRequest) (model.Store, error)
	DeleteStore(userID string, storeID string) error
}

func NewStoreService(client *ApiClient) *StoreService {
	return &StoreService{
		client: client,
	}
}

// CreateStore Create a store
func (s *StoreService) CreateStore(userID string, body model.CreateStoreRequest) (model.Store, error) {
	var (
		path = "%s/users/%s/stores"
	)

	postBody, err := json.Marshal(body)
	if err != nil {
		return model.Store{}, err
	}

	url := fmt.Sprintf(path, s.client.BaseURL, userID)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(postBody))
	if err != nil {
		return model.Store{}, err
	}

	// request url
	r, err := s.client.callAPI(req)
	if err != nil {
		return model.Store{}, err
	}
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response model.Store
	json.Unmarshal(b, &response)

	return response, nil

}

// GetStore Get a store by ID
func (s *StoreService) GetStore(storeID string) (model.GetStoreResponse, error) {

	var (
		path = "/stores"
	)

	url := fmt.Sprintf("%s%s/%s", s.client.BaseURL, path, storeID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.GetStoreResponse{}, err
	}

	// request url
	r, err := s.client.callAPI(req)
	if err != nil {
		return model.GetStoreResponse{}, err
	}
	defer r.Body.Close()

	var response model.GetStoreResponse
	s.client.DeserializeBody(r.Body, &response)

	return response, nil

}

// GetStore By External ID Get a store by external ID
func (s *StoreService) GetStoreByExternalID(userID string, externalID string) (model.GetStoresResponse, error) {

	url := fmt.Sprintf("%susers/%s/stores/search?external_id=%s", s.client.BaseURL, userID, externalID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.GetStoresResponse{}, err
	}

	// request url
	r, err := s.client.callAPI(req)
	if err != nil {
		return model.GetStoresResponse{}, err
	}

	defer r.Body.Close()
	var response model.GetStoresResponse
	s.client.DeserializeBody(r.Body, &response)

	return response, nil

}

// Get Stores Get all stores paginated
func (s *StoreService) GetStores(userID string, limit int) (model.GetStoresResponse, error) {

	var (
		path = "/stores"
	)

	url := fmt.Sprintf("%susers/%s/%s/search?limit=%d", s.client.BaseURL, userID, path, limit)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.GetStoresResponse{}, err
	}

	// request url
	r, err := s.client.callAPI(req)
	if err != nil {
		return model.GetStoresResponse{}, err
	}

	var response model.GetStoresResponse
	err = s.client.DeserializeBody(r.Body, &response)
	if err != nil {
		return model.GetStoresResponse{}, err
	}
	defer r.Body.Close()

	return response, nil

}

func (s *StoreService) UpdateStore(userID string, storeID string, body model.UpdateStoreRequest) (model.Store, error) {

	var (
		path = "%s/users/%s/stores/%s"
	)

	url := fmt.Sprintf(path, s.client.BaseURL, userID, storeID)

	postBody, err := json.Marshal(body)
	if err != nil {
		return model.Store{}, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(postBody))
	if err != nil {
		return model.Store{}, err
	}

	r, err := s.client.callAPI(req)
	if err != nil {
		return model.Store{}, err
	}

	defer r.Body.Close()
	var response model.Store
	s.client.DeserializeBody(r.Body, &response)

	return response, nil

}

func (s *StoreService) DeleteStore(userID string, storeID string) error {

	var (
		path = "%s/users/%s/stores/%s"
	)

	url := fmt.Sprintf(path, s.client.BaseURL, userID, storeID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	// request url
	_, err = s.client.callAPI(req)
	if err != nil {
		return err
	}

	return nil

}
