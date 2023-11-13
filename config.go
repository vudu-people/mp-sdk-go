package mpsdkgo

import "net/http"

type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

type ServerConfigurations []ServerConfiguration

type Config struct {
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	HTTPClient    *http.Client
}

// ServerVariable stores the information about a server variab
// NewConfiguration returns a new Configuration object
func NewConfiguration(options ...func(*Config)) *Config {
	cfg := &Config{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/0.0.4/go",
		Debug:         false,
		Scheme:        "https",
		Host:          "api.mercadopago.com",
	}

	for _, o := range options {
		o(cfg)
	}

	return cfg
}

type ConfigOptions struct {
	Token string
	Debug bool
}

type ConfigOption func(*Config)

func WithDebug(debug bool) ConfigOption {
	return func(o *Config) {
		o.Debug = debug
	}
}

func WithToken(token string) ConfigOption {
	return func(o *Config) {
		o.DefaultHeader["Authorization"] = "Bearer " + token
	}
}
