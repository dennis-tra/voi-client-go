package voi

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_SendLoginMail_HappyPath(t *testing.T) {

	testEmail := "test-email"

	client, teardown := NewTestClientServer(t, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/auth/sso/send", r.URL.Path)

		bodyBytes, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)

		body := map[string]string{}
		err = json.Unmarshal(bodyBytes, &body)
		require.NoError(t, err)

		assert.Equal(t, testEmail, body["email"])

		rw.WriteHeader(http.StatusOK)
		_, err = fmt.Fprint(rw, "") // yes, the API returns an empty string
		require.NoError(t, err)
	}))
	defer teardown(t)

	err := client.SendLoginMail(testEmail)
	assert.NoError(t, err)
}

func TestClient_SendLoginMail_BadRequest(t *testing.T) {

	client, teardown := NewTestClientServer(t, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(rw, "") // yes, the API returns an empty string
		require.NoError(t, err)
	}))
	defer teardown(t)

	err := client.SendLoginMail("test-email")

	assert.IsType(t, &ErrorVoiBadRequest{}, err)
}

func TestClient_SendLoginMail_AlreadyRegistered(t *testing.T) {

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

func TestClient_SendLoginMail_UnexpectedResponseCode(t *testing.T) {

	client, teardown := NewTestClientServer(t, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusBadGateway)
	}))
	defer teardown(t)

	err := client.SendLoginMail("test-email")
	assert.IsType(t, &ErrorVoiUnexpectedResponseCode{}, err)
}
