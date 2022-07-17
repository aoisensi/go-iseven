package iseven

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{client: httpClient}
}

func (c *Client) IsEven(n int) (*Response, error) {
	url := "https://api.isevenapi.xyz/api/iseven/" + strconv.Itoa(n)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("iseven: %v", err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		err := new(ErrorResponse)
		if err := json.Unmarshal(data, err); err != nil {
			return nil, fmt.Errorf("iseven: %v", err)
		}
		return nil, err
	}
	response := new(Response)
	json.Unmarshal(data, response)
	return response, nil
}
