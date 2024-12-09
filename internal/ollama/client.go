package ollama

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/samber/lo"
	"github.com/smbl64/wiz/internal/config"
)

type Client struct {
	baseURL string
}

func Default() *Client {
	return &Client{
		baseURL: config.Current().OllamaAPIBase,
	}
}

func (c *Client) ListModels(ctx context.Context) ([]string, error) {

	fullURL, err := url.JoinPath(c.baseURL, "/api/tags/")
	if err != nil {
		return nil, fmt.Errorf("cannot make url: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot make request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()

	var respBody ListResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to ummarshal body: %v", err)
	}

	return lo.Map(respBody.Models, func(m ListModelResponse, _ int) string {
		return m.Name
	}), nil

}

type ListResponse struct {
	Models []ListModelResponse `json:"models"`
}

// ListModelResponse is a single model description in [ListResponse].
type ListModelResponse struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}
