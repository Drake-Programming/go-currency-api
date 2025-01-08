package api

import (
	"encoding/json"
	"net/http"
)

// Balance Params
type BankBalanceParams struct {
	Username string
}

// Balance Response
type BankBalanceResponse struct {
	// Success Code
	Code int

	// Account Balance
	Balance int64
}

type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}

func writeError(writer http.ResponseWriter, message string, code int) {
	response := Error{
		Code:    code,
		Message: message,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	json.NewEncoder(writer).Encode(response)
}

var (
	RequestErrorHandler = func(writer http.ResponseWriter, err error) {
		writeError(writer, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(writer http.ResponseWriter) {
		writeError(writer, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
