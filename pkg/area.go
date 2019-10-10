package voi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
)

func (c *Client) GetAreas(zone int) (*GetAreasResponse, error) {

	endpoint := baseUrlApi
	endpoint.Path = path.Join(baseUrlApi.Path, "zones/zone/", strconv.Itoa(zone), "/areas")

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

	r := &GetAreasResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, err
}

type GetAreasResponse struct {
	Type     string           `json:"type"`
	Features []FeatureElement `json:"features"`
	CRS      CRS              `json:"crs"`
}

type CRS struct {
	Properties CRSProperties `json:"properties"`
	Type       string        `json:"type"`
}

type CRSProperties struct {
	Name string `json:"name"`
}

type FeatureElement struct {
	ID         string            `json:"id"`
	Type       FeatureType       `json:"type"`
	Geometry   Geometry          `json:"geometry"`
	Properties FeatureProperties `json:"properties"`
}

type Geometry struct {
	Type        GeometryType    `json:"type"`
	Coordinates [][][][]float64 `json:"coordinates"`
}

type FeatureProperties struct {
	AreaType    AreaType `json:"area_type"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Priority    int64    `json:"priority"`
	Rules       Rules    `json:"rules"`
	ZoneID      string   `json:"zone_id"`
}

type Rules struct {
	IsBountyCharge *IsBountyCharge `json:"isBounty_charge,omitempty"`
}

type IsBountyCharge struct {
	Battery   int64    `json:"battery"`
	Days      []string `json:"days"`
	EndTime   string   `json:"end_time"`
	StartTime string   `json:"start_time"`
}

type GeometryType string

const (
	MultiPolygon GeometryType = "MultiPolygon"
)

type AreaType string

const (
	Bounty     AreaType = "bounty"
	NoParking  AreaType = "no-parking"
	Operations AreaType = "operations"
)

type FeatureType string

const (
	Feature FeatureType = "Feature"
)
