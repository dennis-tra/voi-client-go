package voi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
)

// SignUp behaves like the following:
// If a valid, not yet registered email address is provided this endpoint will return a JWT right away
// with the following content:
//
// Header:
//	{
//		"alg": "ES256",
//		"typ": "JWT"
//	}
//
// Payload:
//	{
//		"exp": 1577731695,
//		"jti": "0",
//		"iat": 1569955695,
//		"iss": "auth.api.voiapp.io",
//		"nbf": 1569955694,
//		"userId": "55d07dd9-4a95-44b2-afbd-b7211dd0d223",
//		"UserID": "55d07dd9-4a95-44b2-afbd-b7211dd0d223",
//		"Verified": false,
//		"verified": false
//	}
//
// It is interesting that the custom fields are twice in here :-P
//
// * If a valid and registered email is provided this endpoint returns a 412 Precondition Failed which is not really
//   what one would expect. In the app this seems to indicate to call the `SendLoginMail` endpoint.
// * If an invalid email address is provided this endpoint returns a 400 Bad Request
func (c *Client) SignUp(email string) (*SignUpResponse, error) {

	endpoint := baseUrlApi
	endpoint.Path = path.Join(endpoint.Path, "/auth/sso/signup")

	payload := map[string]string{
		"email": email,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	c.DefaultHeaders.fill(req)

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
	case http.StatusCreated:
		break
	case http.StatusBadRequest:
		return nil, &ErrorVoiBadRequest{Response: resp}
	case http.StatusPreconditionFailed:
		return nil, &ErrorUserAlreadySignedUp{Response: resp}
	default:
		return nil, &ErrorVoiUnexpectedResponseCode{Response: resp}
	}

	r := &SignUpResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, err
}

// SignUpResponse contains the token that was returned from the sign up request
type SignUpResponse struct {
	Token string `json:"authenticationToken"`
}
