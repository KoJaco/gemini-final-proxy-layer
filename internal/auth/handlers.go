package auth

import (
	"encoding/json"

	"net/http"
)

func ClerkPublishableKeyHandler(key string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"publishableKey": "` + key + `"}`))
}

// AuthStatusHandler verifies the session and auth status.
func AuthStatusHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.Header.Get("Authorization")

	if sessionToken == "" {
		http.Error(w, "Unauthorized: No session token provided", http.StatusUnauthorized)
		return
	}

	// verify session
	ctx := GetContext()
	_, verifiedSession, err := VerifySession(ctx, sessionToken) // are we even using ctx again?

	if err != nil {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return
	}

	// return auth status
	authStatus := map[string]interface{}{
		"status":     "authenticated",
		"session_id": verifiedSession.ID,                       // may need to be on the RegisteredClaims
		"user_id":    verifiedSession.RegisteredClaims.Subject, // is this the userId? standard JWT claims
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authStatus)

}

// CreateUserHandler handles the user creation request from the Companion's frontend.
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Emails   []string `json:"email"`
		Password string   `json:"password"`
	}

	ctx := GetContext()

	// Decode the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the user via Clerk API
	newUser, err := CreateUser(ctx, req.Emails, req.Password)

	if err != nil {
		http.Error(w, "Could not create user", http.StatusNotAcceptable)
		return
	}

	if err != nil {
		http.Error(w, "Could not create user", http.StatusNotAcceptable)
		return
	}

	// Return the created user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}
