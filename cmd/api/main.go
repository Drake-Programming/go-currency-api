package main

import (
	"fmt"
	"net/http"

	"github.com/Drake-Programming/go-currency-api/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var route *chi.Mux = chi.NewRouter()
	handlers.Handler(route)

	fmt.Println("Starting GO API service...")
	err := http.ListenAndServe("localhost:8000", route)
	if err != nil {
		log.Error(err)
	}

}
