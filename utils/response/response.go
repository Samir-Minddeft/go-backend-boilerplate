package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Samir-Minddeft/go-backend-boilerplate/utils/types"
	"github.com/go-playground/validator/v10"
)

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) types.Response {
	return types.Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) types.Response {
	var errMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("%s : %s is required", err.Field(), err.Field()))
		default:
			errMsg = append(errMsg, fmt.Sprintf("%s : %s is invalid", err.Field(), err.Field()))

		}
	}

	return types.Response{
		Status: StatusError,
		Error:  strings.Join(errMsg, ", "),
	}
}

func ValidationErrors(errs []string) types.Response {
	return types.Response{
		Status: StatusError,
		Error:  strings.Join(errs, ", "),
	}
}
