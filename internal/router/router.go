package router

import (
	"companion/internal/auth"
	"companion/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// public
	router.HandleFunc("/v1/auth/signup", auth.CreateUserHandler).Methods("POST")

	// stripe routes
	router.HandleFunc("/v1/stripe/create-setup-intent", handlers.CreateSetupIntent).Methods("POST")
	router.HandleFunc("/v1/stripe/session-status", handlers.CreatePaymentIntent).Methods("GET")

	return router
}
