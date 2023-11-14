package mpsdkgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vudu-people/mp-sdk-go/model"
)

type PosService struct {
	client *ApiClient
}

type IPosService interface {
	CreatePos(request model.CreatePosRequest) (model.Pos, error)
	GetPos(posID string) (model.Pos, error)
	GetPosByExternalID(externalID string) (model.GetAllPosResponse, error)
	GetAllPos(limit int) (model.GetAllPosResponse, error)
	UpdatePos(posID string, request model.CreatePosRequest) (model.Pos, error)
	DeletePos(posID string) error
}

func NewPosService(client *ApiClient) *PosService {
	return &PosService{client}
}

func (s *PosService) CreatePos(body model.CreatePosRequest) (model.Pos, error) {
	var (
		path = "%s/pos"
	)

	postBody, err := json.Marshal(body)
	if err != nil {
		return model.Pos{}, err
	}

	url := fmt.Sprintf(path, s.client.BaseURL)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(postBody))
	if err != nil {
		return model.Pos{}, err
	}

	r, err := s.client.callAPI(req)
	if err != nil {
		return model.Pos{}, err
	}

	defer r.Body.Close()

	var response model.Pos
	s.client.DeserializeBody(r.Body, &response)

	return response, nil

}

func (s *PosService) GetPos(posID string) (model.Pos, error) {
	var (
		path = "%s/pos/%s"
	)

	url := fmt.Sprintf(path, s.client.BaseURL, posID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.Pos{}, err
	}

	r, err := s.client.callAPI(req)
	if err != nil {
		return model.Pos{}, err
	}

	defer r.Body.Close()

	var response model.Pos
	s.client.DeserializeBody(r.Body, &response)

	return response, nil
}

func (s *PosService) GetPosByExternalID(externalID string) (model.GetAllPosResponse, error) {
	var (
		path = "%s/pos?external_id=%s"
	)

	url := fmt.Sprintf(path, s.client.BaseURL, externalID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.GetAllPosResponse{}, err
	}

	r, err := s.client.callAPI(req)
	if err != nil {
		return model.GetAllPosResponse{}, err
	}

	defer r.Body.Close()

	var response model.GetAllPosResponse
	s.client.DeserializeBody(r.Body, &response)

	return response, nil
}

func (s *PosService) GetAllPos(limit int) (model.GetAllPosResponse, error) {
	var (
		path = "%s/pos?limit=%d"
	)

	url := fmt.Sprintf(path, s.client.BaseURL, limit)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.GetAllPosResponse{}, err
	}

	r, err := s.client.callAPI(req)
	if err != nil {
		return model.GetAllPosResponse{}, err
	}

	defer r.Body.Close()

	var response model.GetAllPosResponse
	s.client.DeserializeBody(r.Body, &response)

	return response, nil
}

func (s *PosService) UpdatePos(posID string, body model.CreatePosRequest) (model.Pos, error) {

	var (
		path = "%s/pos/%s"
	)

	postBody, err := json.Marshal(body)
	if err != nil {
		return model.Pos{}, err
	}

	url := fmt.Sprintf(path, s.client.BaseURL, posID)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(postBody))
	if err != nil {
		return model.Pos{}, err
	}

	r, err := s.client.callAPI(req)
	if err != nil {
		return model.Pos{}, err
	}

	defer r.Body.Close()

	var response model.Pos
	s.client.DeserializeBody(r.Body, &response)

	return response, nil

}

func (s *PosService) DeletePos(posID string) error {
	var (
		path = "%s/pos/%s"
	)

	url := fmt.Sprintf(path, s.client.BaseURL, posID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.callAPI(req)
	if err != nil {
		return err
	}

	return nil
}
