package voi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func (c *Client) GetUser() (*User, error) {

	endpoint := baseUrlApi
	endpoint.Path = path.Join(baseUrlApi.Path, "/user")

	req, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	c.DefaultHeaders.fill(req)
	req.Header.Set("x-access-token", c.AccessToken)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		break
	default:
		return nil, fmt.Errorf("%v", resp)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &User{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, err
}

type User struct {
	AddedAt string `json:"addedAt"`
	Country string `json:"country"`
	Email   string `json:"email"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Roles   Roles  `json:"roles"`
}

type Roles struct {
	Global []string `json:"global"`
	Zones  Zones    `json:"zones"`
}

type Zones struct {
}
