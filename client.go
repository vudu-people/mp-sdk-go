package mpsdkgo

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type ApiClient struct {
	cfg     *Config
	Store   *StoreService
	BaseURL string
}

func NewApiClient(cfg *Config) *ApiClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	client := &ApiClient{
		cfg:     cfg,
		BaseURL: fmt.Sprintf("%s://%s/", cfg.Scheme, cfg.Host),
	}

	client.Store = NewStoreService(client)

	return client

}

func (c *ApiClient) callAPI(request *http.Request) (*http.Response, error) {
	for k, v := range c.cfg.DefaultHeader {
		request.Header.Add(k, v)
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}
