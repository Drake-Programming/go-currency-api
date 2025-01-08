package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Drake-Programming/go-currency-api/api"
	"github.com/Drake-Programming/go-currency-api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetBalance(writer http.ResponseWriter, reader *http.Request) {
	var params = api.BankBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, reader.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(writer)
		return
	}

	var tokenDetails *tools.BankDetails
	tokenDetails = (*database).GetUserBalance(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var response = api.BankBalanceResponse{
		Balance: (*tokenDetails).Balance,
		Code:    http.StatusOK,
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}
}
