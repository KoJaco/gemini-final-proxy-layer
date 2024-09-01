package auth

import (
	"context"
	"log"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/organization"
	"github.com/clerk/clerk-sdk-go/v2/organizationmembership"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

// CreateUser creates a user in clerk
func CreateUser(ctx context.Context, emails []string, password string) (*clerk.User, error) {

	newUser, err := user.Create(ctx, &user.CreateParams{
		EmailAddresses: &emails,
		Password:       &password,
	})

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return newUser, nil
}

func CreateOrganization(ctx context.Context, name string) (*clerk.Organization, error) {

	org, err := organization.Create(ctx, &organization.CreateParams{
		Name: clerk.String(name),
	})
	if err != nil {
		if apiErr, ok := err.(*clerk.APIErrorResponse); ok {
			log.Printf("Error Trace ID: %s, Error: %s", apiErr.TraceID, apiErr.Error())
		}
		return nil, err
	}

	return org, nil
}

func UpdateOrganization(ctx context.Context, orgID, slug string) (*clerk.Organization, error) {

	org, err := organization.Update(ctx, orgID, &organization.UpdateParams{
		Slug: clerk.String(slug),
	})
	if err != nil {
		return nil, err
	}

	return org, nil
}

func ListOrganizationMemberships(ctx context.Context, limit int64) (*clerk.OrganizationMembershipList, error) {

	listParams := organizationmembership.ListParams{}
	listParams.Limit = clerk.Int64(limit)

	memberships, err := organizationmembership.List(ctx, &listParams)
	if err != nil {
		return nil, err
	}

	return memberships, nil
}

func GetUser(ctx context.Context, userID string) (*clerk.User, error) {

	usr, err := user.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
