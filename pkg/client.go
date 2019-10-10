package voi

import (
	"net/http"
)

type Client struct {
	HttpClient     *http.Client
	DefaultHeaders *Headers

	AccessToken  string
	RefreshToken string
}

func NewClient() *Client {
	return &Client{
		HttpClient:     &http.Client{},
		DefaultHeaders: NewDefaultHeaders(),
	}
}

// Authenticate takes the Link or token behind the "Open the app" button in the login email and prepares the client
// for all subsequent requests.
// Valid values for `link` are either:
// * https://link.voiapp.io/ohJ7i9ahr0
// * ohJ7i9ahr0
// This endpoint persists tokens in the client
func (c *Client) Authenticate(link string) error {

	branchIoToken, err := c.VerifyLoginMail(link)
	if err != nil {
		return err
	}

	session, err := c.Session(branchIoToken)
	if err != nil {
		return err
	}

	c.AccessToken = session.AccessToken
	c.RefreshToken = session.AuthenticationToken

	return nil
}
