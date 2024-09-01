package main

import (
	"log"

	"companion/internal/auth"
)

func main() {
	// Initialize Clerk and load environment variables
	auth.InitClerk()

	// Set up the router (assuming you have a router setup in internal/api)
	// router := api.NewRouter()

	// Start the HTTP server
	log.Println("Starting server on :8080")
	// if err := http.ListenAndServe(":8080", router); err != nil {
	// 	log.Fatal(err)
	// }
}
