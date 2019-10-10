package voi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func (c *Client) GetZones(latitude, longitude float32) (*GetZoneResponse, error) {

	endpoint := baseUrlApi
	endpoint.Path = path.Join(baseUrlApi.Path, "/zones")

	v := endpoint.Query()
	v.Set("lat", fmt.Sprintf("%f", latitude))
	v.Set("lng", fmt.Sprintf("%f", longitude))
	endpoint.RawQuery = v.Encode()

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

	r := &GetZoneResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, err
}

type GetZoneResponse struct {
	Zones []Zone `json:"zones"`
}

type Zone struct {
	ZoneID              string        `json:"zone_id"`
	Name                string        `json:"name"`
	Country             string        `json:"country"`
	CompanyName         string        `json:"company_name"`
	CompanyID           string        `json:"company_id"`
	Currency            string        `json:"currency"`
	Vat                 float64       `json:"vat"`
	StartCost           int64         `json:"start_cost"`
	MinuteCost          int64         `json:"minute_cost"`
	FeeOutOfZone        int64         `json:"fee_out_of_zone"`
	CreditsExchangeRate int64         `json:"credits_exchange_rate"`
	Language            string        `json:"language"`
	DefaultLocale       string        `json:"default_locale"`
	Locales             string        `json:"locales"`
	TimeZone            string        `json:"time_zone"`
	MaxSpeed            int64         `json:"max_speed"`
	Boundaries          Boundaries    `json:"boundaries"`
	BountyStateExpr     interface{}   `json:"bounty_state_expr"`
	BountyAmountExpr    string        `json:"bounty_amount_expr"`
	MinutePrice         []MinutePrice `json:"minute_price"`
	IsActive            bool          `json:"is_active"`
}

type Boundaries struct {
	Lo Coordinate `json:"Lo"`
	Hi Coordinate `json:"Hi"`
}

type Coordinate struct {
	Latitude  float64 `json:"Lat"`
	Longitude float64 `json:"Lng"`
}

type MinutePrice struct {
	Day    *string `json:"day,omitempty"`
	Period Period  `json:"period"`
	Amount int64   `json:"amount"`
}

type Period string

const (
	Day   Period = "day"
	Night Period = "night"
)
