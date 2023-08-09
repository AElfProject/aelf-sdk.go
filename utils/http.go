package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	Endpoints         []string
	TimeoutPerRequest time.Duration
}

type httpClient struct {
	endpoints []string
	Client    *http.Client
}

func New(cfg Config) (*httpClient, error) {
	c := &httpClient{
		Client: &http.Client{
			Timeout: cfg.TimeoutPerRequest,
		},
	}

	return c, nil
}

func (c *httpClient) Do(method, url, version, basicAuth string, params map[string]interface{}) (*http.Response, error) {
	var headerKey string
	var body io.Reader

	switch method {
	case "POST":
		if params != nil {
			bytesData, _ := json.Marshal(params)
			body = strings.NewReader(string(bytesData))
		}
		headerKey = "Content-Type"
	case "GET":
		if params != nil {
			strParams := Map2UrlParams(params)
			url += "?" + strParams
		}
		headerKey = "Accept"
	default:
		return &http.Response{}, errors.New("unexpected http method")
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return &http.Response{}, err
	}

	if version != "" {
		request.Header.Set(headerKey, "application/json;v="+version)
	} else {
		request.Header.Set(headerKey, "application/json")
	}
	if basicAuth != "" {
		request.Header.Set("Authorization", basicAuth)
	}

	resp, err := c.Client.Do(request)
	return resp, err
}
