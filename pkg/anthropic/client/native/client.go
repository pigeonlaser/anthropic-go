package native

import (
	"net/http"

	"github.com/madebywelch/anthropic-go/v3/pkg/anthropic"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

type Config struct {
	APIKey  string
	BaseURL string
}

func MakeClient(cfg Config) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, anthropic.ErrAnthropicApiKeyRequired
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.anthropic.com"
	}

	return &Client{
		httpClient: &http.Client{},
		apiKey:     cfg.APIKey,
		baseURL:    cfg.BaseURL,
	}, nil
}
