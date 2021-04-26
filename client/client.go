package client

import (
	"log"
	"net/http"
	"net/url"
)

var httpClient = http.DefaultClient

type Client struct {
	http.Client
	token string
	url   string
}

func New(api string, token string) Client {
	c := Client{
		token: token,
		url:   api,
	}
	return c
}

func (c Client) Ping() bool {
	url, err := url.Parse(c.url)
	if err != nil {
		log.Fatalf("api scheme format error '%s' %s", c.url, err)
	}
	req := http.Request{
		Method: "GET",
		URL:    url,
	}
	req.Header.Add("PRIVATE-TOKEN", c.token)
	c.Do(&req)
	return true
}
