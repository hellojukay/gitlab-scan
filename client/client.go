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
	url, err := url.Parse(c.url + "/projects")
	if err != nil {
		log.Fatalf("api scheme format error '%s' %s", c.url, err)
	}
	req := http.Request{
		Method: "GET",
		URL:    url,
		Header: http.Header{},
	}
	req.Header.Set("PRIVATE-TOKEN", c.token)
	res, err := c.Do(&req)
	if err != nil {
		log.Printf("request %s failed, %s", c.url, err)
		return false
	}
	if res.StatusCode == 200 {
		return true
	} else {
		log.Printf("%s %d", c.url, res.StatusCode)
	}
	return false
}
