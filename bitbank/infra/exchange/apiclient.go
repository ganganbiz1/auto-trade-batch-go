package exchange

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/config"
	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/domain/entity"
)

const (
	baseURL = "https://api.bitbank.cc"
)

type APIClient struct {
	config     *config.APIConfig
	httpClient *http.Client
}

func NewAPIClient(
	config *config.APIConfig,
	httpClient *http.Client,
) *APIClient {
	return &APIClient{
		config:     config,
		httpClient: httpClient,
	}
}

func (c *APIClient) sendGetRequest(ctx context.Context, path string, query map[string]string) ([]byte, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	apiURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	endpoint := u.ResolveReference(apiURL).String()

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	apiInfo, err := entity.NewAPIInfo("GET", path, query, nil)
	if err != nil {
		return nil, err
	}

	auth := entity.NewAuth(c.config, apiInfo)

	header := auth.CreateRequestHeader()
	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *APIClient) sendPostRequest(ctx context.Context, path string, data []byte) ([]byte, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	apiURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	endpoint := u.ResolveReference(apiURL).String()

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	apiInfo, err := entity.NewAPIInfo("POST", path, nil, data)
	if err != nil {
		return nil, err
	}
	auth := entity.NewAuth(c.config, apiInfo)

	header := auth.CreateRequestHeader()
	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *APIClient) GetAsset(ctx context.Context) ([]*entity.Asset, error) {
	_, err := c.sendGetRequest(ctx, "/v1/user/assets", nil)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return nil, nil
}

func (c *APIClient) Order(ctx context.Context, order entity.Order) error {
	log.Println("order:", order)
	return nil
}
