package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func main() {
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		// Here you can check the request, for example:
		if r.URL.Path != "/status" {
			panic("failed test!")
		}

		w.Write([]byte(`{"status":"OK!"}`))
	})
	defer ts.Close()

	c, err := NewClient(BaseURL(ts.URL))
	if err != nil {
		panic(err)
	}

	sr, err := c.GetStatus("/status")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", sr)
}

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
}

type Option func(*Client) error

func BaseURL(rawurl string) Option {
	return func(c *Client) error {
		baseURL, err := url.Parse(rawurl)
		if err != nil {
			return err
		}

		c.baseURL = baseURL

		return nil
	}
}

func NewClient(options ...Option) (*Client, error) {
	c := &Client{httpClient: &http.Client{}}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) GetStatus(path string) (*StatusResponse, error) {
	u := c.baseURL.ResolveReference(&url.URL{Path: path})

	resp, err := c.httpClient.Get(u.String())
	if err != nil {
		return nil, err
	}

	var sr StatusResponse

	return &sr, json.NewDecoder(resp.Body).Decode(&sr)
}

func testServer(hf http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(hf)
}
