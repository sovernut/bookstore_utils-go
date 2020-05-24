package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("message error", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "message error", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])

}

func TestNewRestError(t *testing.T) {
	err := NewRestError("message error", http.StatusInternalServerError, "internal error", nil)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "message error", err.Message)
	assert.EqualValues(t, "internal error", err.Error)
}

func TestNewBadRequestError(t *testing.T) {

}

func TestNewNotFoundError(t *testing.T) {

}

func TestNewError(t *testing.T) {

}
