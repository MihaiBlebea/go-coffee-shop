package server

import (
	"fmt"
	"log"

	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const prefix = "/api/v1/"

func runServer(handler Handler, logger Logger) {
	r := mux.NewRouter().
		PathPrefix(prefix).
		Subrouter()

	r.Handle("/health-check", handler.HealthEndpoint()).
		Methods("GET")

	r.Handle("/order", handler.OrderEndpoint()).
		Methods(http.MethodPost)

	r.Handle("/orders", handler.OrdersEndpoint()).
		Methods(http.MethodGet)

	r.Handle("/menu", handler.MenuEndpoint()).
		Methods(http.MethodGet)

	r.Use(loggerMiddleware(logger))

	srv := &http.Server{
		Handler:      cors.Default().Handler(r),
		Addr:         fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT_ORDER")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
