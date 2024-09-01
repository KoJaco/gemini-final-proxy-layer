package auth

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/joho/godotenv"
)

func InitClerk() {

	// Print the current working directory
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	err := godotenv.Load("../../.env.local")

	if err != nil {
		log.Fatalf("Error loading .env.local file")
	}

	apiKey := os.Getenv("CLERK_SECRET_KEY")

	if apiKey == "" {
		log.Fatal("CLERK_API_KEY environment variable not set")
	}

	clerk.SetKey(apiKey)

}

func GetContext() context.Context {
	return context.Background()
}
