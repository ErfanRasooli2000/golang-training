package errs

import (
	"net/http"
)

type AppError struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NotFoundHttpError(message string) *AppError {

	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
		Status:  false,
	}

}

func ServerError(message string) *AppError {

	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Status:  false,
	}
}

func HandleErrorResponse(writer http.ResponseWriter, error *AppError) {

	writer.Header().Add("content-type", "application/json")

}
