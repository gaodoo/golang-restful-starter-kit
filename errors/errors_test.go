package errors

import (
	"net/http"
	"testing"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {
	assert.Equal(t, http.StatusInternalServerError, InternalServerError(errors.New("")).Status)
}

func TestInvalidData(t *testing.T) {
	err := InvalidData(validation.Errors{
		"abc": errors.New("1"),
		"xyz": errors.New("2"),
	})
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.NotNil(t, err.Details)
}

func TestNotFound(t *testing.T) {
	assert.Equal(t, http.StatusNotFound, NotFound("abc").Status)
}
