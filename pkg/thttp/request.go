package thttp

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

var GET_METHOD = "GET"

type Client struct {
	cli *http.Client
}

func GetClient() *Client {
	return &Client{
		cli: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) SetTimeout(sec int) {
	c.cli.Timeout = time.Second * time.Duration(sec)
}

func (c *Client) Do(ctx context.Context, method, url string, header map[string]string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	if err != nil {
		return "", err
	}
	resp, err := c.cli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
