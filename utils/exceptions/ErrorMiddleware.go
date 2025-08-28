package exceptions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Status  int  `json:"statusCode"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type NotFoundError struct {
	Details string
    Resource string
    ID       interface{}
}

type BadRequestError struct {
	Details string
    Field   string
    Message string
}

type InternalServerError struct {
    Message   string
}


func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s not found: %v", e.Resource, e.ID)
}

func (e *InternalServerError) Error() string {
    return e.Message
}

func (e *BadRequestError) Error() string {
    if e.Field != "" {
        return fmt.Sprintf("invalid field %s: %s", e.Field, e.Message)
    }
    return e.Message
}

func NewNotFound(resource, details string, id interface{}) error {
    return &NotFoundError{Resource: resource, ID: id, Details: details}
}

func NewBadRequest(field, message, details string) error {
    return &BadRequestError{Field: field, Message: message, Details: details}
}

func NewBadRequestSimple(message, details string) error {
    return &BadRequestError{Message: message, Details: details}
}

func NewInternalServer(message string) error {
    return &InternalServerError{Message: message}
}

func WriteError(w http.ResponseWriter, e ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(e)
}

func GetNotFoundError(message, details string) ErrorResponse {
	return ErrorResponse{
		Status: 404,
		Message: message,
		Details: details,
	}
}

func GetBadRequestError(message, details string) ErrorResponse {
	return ErrorResponse{
		Status: 400,
		Message: message,
		Details: details,
	}
}