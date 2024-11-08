package config

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(messsage, err string, code int) *RestErr {
	return &RestErr{
		Message: messsage,
		Err:     err,
		Code:    code,
	}
}

func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewNotFoundErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewInternaErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewConflictErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusConflict,
	}
}
