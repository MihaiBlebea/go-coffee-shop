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

	r.Handle("/user", handler.AuthenticateEndpoint()).
		Methods(http.MethodPost)

	r.Use(loggerMiddleware(logger))

	srv := &http.Server{
		Handler:      cors.Default().Handler(r),
		Addr:         fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT_ACCOUNT")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
