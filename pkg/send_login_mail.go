package voi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path"
)

// SendLoginMail sends an email to the provided email address with a log in link.
// This link contains a token that should be used to validate the login.
// To validate the login call `Login`.
func (c *Client) SendLoginMail(email string) error {

	endpoint := baseUrlApi
	endpoint.Path = path.Join(endpoint.Path, "/auth/sso/send")

	payload := map[string]string{
		"email": email,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint.String(), bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	c.DefaultHeaders.fill(req)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return &ErrorVoiBadRequest{Response: resp}
	default:
		return &ErrorVoiUnexpectedResponseCode{Response: resp}
	}
}
