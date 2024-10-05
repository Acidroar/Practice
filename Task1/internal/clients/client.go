package clients

import (
	"net/http"
)

type Client struct {
	*http.Client
}

func New() *Client {
	return &Client{
		Client: &http.Client{},
	}
}
