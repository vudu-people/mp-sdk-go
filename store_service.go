package mpsdkgo

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

func NewStoreService(client *ApiClient) *StoreService {
	return &StoreService{
		client: client,
	}
}

// CreateStore Create a store
func (s *StoreService) CreateStore(body model.CreateStoreRequest) (model.CreateStoreResponse, error) {
	var (
		path = "/stores"
	)

	postBody, err := json.Marshal(body)
	if err != nil {
		return model.CreateStoreResponse{}, err
	}

	url := fmt.Sprintf("%s%s", s.client.BaseURL, path)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(postBody))
	if err != nil {
		return model.CreateStoreResponse{}, err
	}

	// request url
	r, err := s.client.callAPI(req)
	if err != nil {
		return model.CreateStoreResponse{}, err
	}
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response model.CreateStoreResponse
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
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var response model.GetStoreResponse
	json.Unmarshal(b, &response)

	return response, nil

}

// GetStore By External ID Get a store by external ID
func (s *StoreService) GetStoreByExternalID(userID string, externalID string) (model.GetStoreResponse, error) {

	var (
		path = "/stores/search?external_id="
	)

	url := fmt.Sprintf("%susers/%s/%s/search?external_id=%s", s.client.BaseURL, userID, path, externalID)

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
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var response model.GetStoreResponse
	json.Unmarshal(b, &response)

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

	// read response body
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response model.GetStoresResponse
	json.Unmarshal(b, &response)

	return response, nil

}
