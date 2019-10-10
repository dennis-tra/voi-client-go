package voi

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func NewTestClientServer(t *testing.T, handler http.Handler) (*Client, func(*testing.T)) {

	ts := httptest.NewServer(handler)

	newBaseUrlApi, err := url.Parse(ts.URL)
	require.NoError(t, err)

	baseUrlApi = *newBaseUrlApi

	teardown := func(t *testing.T) {
		ts.Close()
	}

	return NewClient(), teardown
}

func TestClient_SignUp_HappyPath(t *testing.T) {

	testEmail := "test-email"

	client, teardown := NewTestClientServer(t, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/auth/sso/signup", r.URL.Path)

		bodyBytes, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)

		body := map[string]string{}
		err = json.Unmarshal(bodyBytes, &body)
		require.NoError(t, err)

		assert.Equal(t, testEmail, body["email"])

		exampleResponse := `{"authenticationToken":"test-token"}`

		rw.WriteHeader(http.StatusCreated)
		_, err = fmt.Fprint(rw, exampleResponse)
		require.NoError(t, err)
	}))
	defer teardown(t)

	resp, err := client.SignUp(testEmail)
	assert.NoError(t, err)

	assert.Equal(t, resp.Token, "test-token")
}

func TestClient_SignUp_BadRequest(t *testing.T) {

	client, teardown := NewTestClientServer(t, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(rw, "") // yes, the API returns an empty string
		require.NoError(t, err)
	}))
	defer teardown(t)

	resp, err := client.SignUp("test-email")
	assert.Nil(t, resp)

	assert.IsType(t, &ErrorVoiBadRequest{}, err)
}

func TestClient_SignUp_AlreadyRegistered(t *testing.T) {

	client, teardown := NewTestClientServer(t, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusPreconditionFailed)
		_, err := fmt.Fprint(rw, "") // yes, the API returns an empty string
		require.NoError(t, err)
	}))
	defer teardown(t)

	resp, err := client.SignUp("test-email")
	assert.Nil(t, resp)

	assert.IsType(t, &ErrorUserAlreadySignedUp{}, err)
}

func TestClient_SignUp_UnexpectedResponseCode(t *testing.T) {

	client, teardown := NewTestClientServer(t, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusBadGateway)
	}))
	defer teardown(t)

	resp, err := client.SignUp("test-email")
	assert.Nil(t, resp)

	assert.IsType(t, &ErrorVoiUnexpectedResponseCode{}, err)
}
