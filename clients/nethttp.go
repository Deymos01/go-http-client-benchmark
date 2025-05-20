package clients

import (
	"io"
	"net/http"
)

type HTTPClient struct {
	Client *http.Client
	URL    string
}

func NewHTTPClient(url string) *HTTPClient {
	return &HTTPClient{
		Client: &http.Client{},
		URL:    url,
	}
}

func (c *HTTPClient) MakeGetRequest() error {
	resp, err := c.Client.Get(c.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(io.Discard, resp.Body)
	return err
}
