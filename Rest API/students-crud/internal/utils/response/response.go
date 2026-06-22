package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(err validator.ValidationErrors) Response {
	var errorMessages []string

	for _, fieldErr := range err {
		switch fieldErr.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is required", fieldErr.Field()))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is invalid: %s", fieldErr.Field(), fieldErr.Error()))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errorMessages, "; "),
	}
}
