package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	v4 "github.com/hellojukay/gitlab/gitlab-scan/v4"
)

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
		Method: http.MethodGet,
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
		defer res.Body.Close()
		return true
	} else {
		log.Printf("can not auth  %s reponse with http code %d", c.url, res.StatusCode)
	}
	return false
}

// Group gitlab group information
func (c Client) Group(id int64) (*v4.Group, error) {
	u := c.url + fmt.Sprintf("/groups/%d", id)
	api, err := url.Parse(u)
	if err != nil {
		log.Printf("prase url %s error %s,check please", u, err)
		return nil, err
	}
	req := http.Request{
		Method: http.MethodGet,
		URL:    api,
		Header: http.Header{},
	}
	req.Header.Set("PRIVATE-TOKEN", c.token)
	res, err := c.Do(&req)
	if err != nil {
		log.Printf("request gitlab grop detail failed, %s  %s", api, err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("can not read any from gitlab %s, %s", api, err)
		return nil, err
	}
	var group v4.Group
	if err := json.Unmarshal(body, &group); err != nil {
		log.Printf("gitlab response data may not a json string,%s", string(body))
		return nil, err
	}
	return &group, nil
}

func (c Client) Projects(g v4.Group) ([]v4.Project, error) {
	u := c.url + fmt.Sprintf("/groups/%d/projects", g.ID)
	api, err := url.Parse(u)
	if err != nil {
		log.Printf("prase url %s error %s,check please", u, err)
		return nil, err
	}
	req := http.Request{
		Method: http.MethodGet,
		URL:    api,
		Header: http.Header{},
	}
	req.Header.Set("PRIVATE-TOKEN", c.token)
	res, err := c.Do(&req)
	if err != nil {
		log.Printf("request gitlab grop detail failed, %s  %s", api, err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("can not read any from gitlab %s, %s", api, err)
		return nil, err
	}
	var projects []v4.Project
	if err := json.Unmarshal(body, &projects); err != nil {
		log.Printf("gitlab response data may not a json string,%s", string(body))
		return nil, err
	}
	return projects, nil
}

func (c Client) Branches(project v4.Project) ([]v4.Branch, error) {
	u := c.url + fmt.Sprintf("/projects/%d/repository/branches", project.ID)
	api, err := url.Parse(u)
	if err != nil {
		log.Printf("prase url %s error %s,check please", u, err)
		return nil, err
	}
	req := http.Request{
		Method: http.MethodGet,
		URL:    api,
		Header: http.Header{},
	}
	req.Header.Set("PRIVATE-TOKEN", c.token)
	res, err := c.Do(&req)
	if err != nil {
		log.Printf("request gitlab group detail failed, %s  %s", api, err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("can not read any from gitlab %s, %s", api, err)
		return nil, err
	}
	var branches []v4.Branch
	if err := json.Unmarshal(body, &branches); err != nil {
		log.Printf("gitlab response data may not a json string,%s", string(body))
		return nil, err
	}
	return branches, nil
}

func (c Client) Subgroups(group v4.Group) ([]v4.Group, error) {
	u := c.url + fmt.Sprintf("/groups/%d/subgroups", group.ID)
	api, err := url.Parse(u)
	if err != nil {
		log.Printf("prase url %s error %s,check please", u, err)
		return nil, err
	}
	req := http.Request{
		Method: http.MethodGet,
		URL:    api,
		Header: http.Header{},
	}
	req.Header.Set("PRIVATE-TOKEN", c.token)
	res, err := c.Do(&req)
	if err != nil {
		log.Printf("request gitlab grop detail failed, %s  %s", api, err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("can not read any from gitlab %s, %s", api, err)
		return nil, err
	}
	var groups []v4.Group
	if err := json.Unmarshal(body, &groups); err != nil {
		log.Printf("gitlab response data may not a json string,%s", string(body))
		return nil, err
	}
	return groups, nil
}
