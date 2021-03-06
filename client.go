package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
)

const apiURL = "https://api.coingecko.com/api/v3"

// Client models a client to consume the Polygon Cloud API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// Error represents an Polygon API error
type Error struct {
	ErrorMessage string `json:"error"`
}

// ClientOption applies an option to the client.
type ClientOption func(*Client)

// Error implements the error interface
func (e Error) Error() string {
	return e.ErrorMessage
}

// NewClient creates a client with the given authorization token.
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{Timeout: time.Second * 60},
	}

	// apply options
	for _, applyOption := range options {
		applyOption(client)
	}

	// set default values
	if client.baseURL == "" {
		client.baseURL = apiURL
	}

	return client
}

// WithHTTPClient sets the http.Client for a new IEX Client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

// WithSecureHTTPClient sets a secure http.Client for a new IEX Client
func WithSecureHTTPClient() ClientOption {
	return func(client *Client) {
		client.httpClient = &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			}}
	}
}

// WithBaseURL sets the baseURL for a new IEX Client
func WithBaseURL(baseURL string) ClientOption {
	return func(client *Client) {
		client.baseURL = baseURL
	}
}

// GetJSON gets the JSON data from the given endpoint.
func (c *Client) GetJSON(ctx context.Context, endpoint string, v interface{}) error {
	u, err := c.url(endpoint, map[string]string{})
	if err != nil {
		return err
	}
	return c.FetchURLToJSON(ctx, u, v)
}

// GetJSONWithQueryParams gets the JSON data from the given endpoint with the query parameters attached.
func (c *Client) GetJSONWithQueryParams(ctx context.Context, endpoint string, queryParams map[string]string, v interface{}) error {
	u, err := c.url(endpoint, queryParams)
	if err != nil {
		return err
	}
	return c.FetchURLToJSON(ctx, u, v)
}

// Fetches JSON content from the given URL and unmarshals it into `v`.
func (c *Client) FetchURLToJSON(ctx context.Context, u *url.URL, v interface{}) error {
	data, err := c.getBytes(ctx, u.String())
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// GetJSONWithoutToken gets the JSON data from the given endpoint without
// adding a token to the URL.
func (c *Client) GetJSONWithoutToken(ctx context.Context, endpoint string, v interface{}) error {
	u, err := c.url(endpoint, nil)
	if err != nil {
		return err
	}
	return c.FetchURLToJSON(ctx, u, v)
}

// GetBytes gets the data from the given endpoint.
func (c *Client) GetBytes(ctx context.Context, endpoint string) ([]byte, error) {
	u, err := c.url(endpoint, map[string]string{})
	if err != nil {
		return nil, err
	}
	return c.getBytes(ctx, u.String())
}

// GetFloat64 gets the number from the given endpoint.
func (c *Client) GetFloat64(ctx context.Context, endpoint string) (float64, error) {
	b, err := c.GetBytes(ctx, endpoint)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(string(b), 64)
}

func (c *Client) getBytes(ctx context.Context, address string) ([]byte, error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	// Even if GET didn't return an error, check the status code to make sure
	// everything was ok.
	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		msg := ""

		if err == nil {
			msg = string(b)
		}

		return []byte{}, Error{ErrorMessage: msg}
	}
	return ioutil.ReadAll(resp.Body)
}

// Returns an URL object that points to the endpoint with optional query parameters.
func (c *Client) url(endpoint string, queryParams map[string]string) (*url.URL, error) {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, err
	}

	if queryParams != nil {
		q := u.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}
	return u, nil
}

func (c Client) endpointWithOpts(endpoint string, opts interface{}) (string, error) {
	if opts == nil {
		return endpoint, nil
	}
	v, err := query.Values(opts)
	if err != nil {
		return "", err
	}
	optParams := v.Encode()
	if optParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, optParams)
	}

	return endpoint, nil
}
