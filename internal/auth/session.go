package auth

import (
	"context"
	"errors"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/go-jose/go-jose/v3/jwt"
)

// VerifySession decodes and validates the session JWT token.
func VerifySession(ctx context.Context, token string) (context.Context, *clerk.SessionClaims, error) {

	// Parse the JWT token without verifying the signature first
	unverifiedToken, err := jwt.ParseSigned(token)
	if err != nil {
		return nil, nil, errors.New("invalid session token")
	}

	// Decode the claims without verifying them
	var claims clerk.SessionClaims
	err = unverifiedToken.UnsafeClaimsWithoutVerification(&claims)
	if err != nil {
		return nil, nil, errors.New("unable to decode session claims")
	}

	// Verify the standard claims with some leeway (2 minutes, may need to adjust this.)
	leeway := time.Minute * 2
	if err := claims.RegisteredClaims.ValidateWithLeeway(time.Now(), leeway); err != nil {
		return nil, nil, errors.New("session token is expired or invalid")
	}

	// Todo: add custom validation for additional roles/permissions
	if !claims.HasRole("admin") {
		return nil, nil, errors.New("user does not have the required role")
	}

	// Embed the claims into the context for further use
	ctx = clerk.ContextWithSessionClaims(ctx, &claims)

	return ctx, &claims, nil
}
