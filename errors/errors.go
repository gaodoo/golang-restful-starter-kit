package errors

import (
	"net/http"
	"sort"

	"github.com/go-ozzo/ozzo-validation"
)

type validationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func InternalServerError(err error) *APIError {
	return NewAPIError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", Params{"error": err.Error()})
}

func NotFound(resource string) *APIError {
	return NewAPIError(http.StatusNotFound, "NOT_FOUND", Params{"resource": resource})
}

func InvalidData(errs validation.Errors) *APIError {
	result := []validationError{}
	fields := []string{}
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		err := errs[field]
		result = append(result, validationError{
			Field: field,
			Error: err.Error(),
		})
	}

	err := NewAPIError(http.StatusBadRequest, "INVALID_DATA", nil)
	err.Details = result

	return err
}
