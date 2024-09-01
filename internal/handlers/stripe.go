package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/setupintent"
)

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
}

// CreateSetupIntent is the initial step of the checkout process, enter card details but do not process --> the user then accepts the review of everything
func CreateSetupIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = "sk_test_51PtHPOByxck9sYSvr9kjmsG0A0J6T2Q49Z1PhsyGU9reFUnmSt4PGZ2x8WRkKUiy5iPNixlRa7CtNXrPNLZDdhAH006OunNI4J"

	params := &stripe.SetupIntentParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}

	si, err := setupintent.New(params)
	if err != nil {
		log.Printf("setupintent.New: %v", err)
		http.Error(w, "Failed to create setup intent", http.StatusInternalServerError)
		return
	}

	writeJSON(w, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: si.ClientSecret,
	})
}

// CreatePaymentIntent handles creating a payment intent to charge the customer, this is after the review of the user details and payment stuff.
func CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = "sk_test_51PtHPOByxck9sYSvr9kjmsG0A0J6T2Q49Z1PhsyGU9reFUnmSt4PGZ2x8WRkKUiy5iPNixlRa7CtNXrPNLZDdhAH006OunNI4J"

	var req struct {
		PaymentMethodID string `json:"paymentMethodId"`
		Amount          int64  `json:"amount"`
		Currency        string `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(req.Amount),
		Currency:      stripe.String(req.Currency),
		PaymentMethod: stripe.String(req.PaymentMethodID),
		Confirm:       stripe.Bool(true),
		OffSession:    stripe.Bool(true),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		log.Printf("paymentintent.New: %v", err)
		http.Error(w, "Failed to create payment intent", http.StatusInternalServerError)
		return
	}

	writeJSON(w, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: pi.ClientSecret,
	})
}
