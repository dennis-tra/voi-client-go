package voi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
)

// Session starts the user session and persists the tokens in the current client
func (c *Client) Session(token string) (*SessionResponse, error) {

	endpoint := baseUrlApi
	endpoint.Path = path.Join(endpoint.Path, "/auth/session")

	payload := map[string]string{
		"authenticationToken": token,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		break
	case http.StatusUnauthorized:
		return nil, &ErrorVoiUnauthroized{}
	default:
		return nil, &ErrorVoiUnexpectedResponseCode{Response: resp}
	}

	r := &SessionResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// SessionResponse is the JSON replied by the start session call
type SessionResponse struct {
	AccessToken         string `json:"AccessToken"`
	AuthenticationToken string `json:"authenticationToken"`
}
