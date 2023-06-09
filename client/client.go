package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/inconshreveable/log15"
)

const (
	baseURL = "https://www.recreation.gov/api/"
)

// Client is an HTTP client that interacts with the recreation.gov API.
type Client struct {
	client *http.Client
	log    log15.Logger
}

func New(l log15.Logger, timeout time.Duration) *Client {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		log: l,
	}
}

func (c *Client) Do(path string, queryParams url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryParams.Encode()

	// Need to spoof the user agent or CloudFront blocks us.
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		c.log.Error("Non-2xx status returned", "status", resp.Status, "response", string(bytes))
		return nil, fmt.Errorf("Non-2xx status: '%s'", resp.Status)
	}

	return resp, nil
}
