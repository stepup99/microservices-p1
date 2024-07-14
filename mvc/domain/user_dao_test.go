package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// initialization

	// Execution
	resp, err := GetUser(0)

	//Validation
	assert.Nil(t, resp, "we are not expecting a user with id 0")
	assert.NotNil(t, err, "we are expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	resp, err := GetUser(123)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, "mudata@email.com", resp.Email)
}
