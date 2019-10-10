package voi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
)

type Status string

const (
	StatusReady     Status = "ready"
	StatusRiding    Status = "riding"
	StatusBounty    Status = "bounty"
	StatusLost      Status = "lost"
	StatusHome      Status = "home"
	StatusCollected Status = "collected"
)

func (c *Client) GetScooters(zone int, status Status) (*Scooters, error) {

	endpoint := baseUrlApi
	endpoint.Path = path.Join(baseUrlApi.Path, "/vehicles/zone/", strconv.Itoa(zone), string(status))

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

	r := &Scooters{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, err
}

type Scooters []Scooter

type Scooter struct {
	Added              string      `json:"added"`
	Battery            int64       `json:"battery"`
	Bounty             int64       `json:"bounty"`
	ID                 string      `json:"id"`
	Location           []float64   `json:"location"`
	Locked             bool        `json:"locked"`
	Mileage            int64       `json:"mileage"`
	ModelSpecification interface{} `json:"model_specification"`
	Name               string      `json:"name"`
	RegistrationPlate  string      `json:"registration_plate"`
	Serial             interface{} `json:"serial"`
	Short              string      `json:"short"`
	Status             string      `json:"status"`
	Type               string      `json:"type"`
	Updated            string      `json:"updated"`
	Zone               int64       `json:"zone"`
}
