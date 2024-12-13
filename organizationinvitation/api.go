// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package organizationinvitation

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v3"
)

// Create creates and sends an invitation to join an organization.
func Create(ctx context.Context, params *CreateParams) (*clerk.OrganizationInvitation, error) {
	return getClient().Create(ctx, params)
}

// List returns a list of organization invitations
func List(ctx context.Context, params *ListParams) (*clerk.OrganizationInvitationList, error) {
	return getClient().List(ctx, params)
}

// Get retrieves the detail for an organization invitation.
func Get(ctx context.Context, params *GetParams) (*clerk.OrganizationInvitation, error) {
	return getClient().Get(ctx, params)
}

// Revoke marks the organization invitation as revoked.
func Revoke(ctx context.Context, params *RevokeParams) (*clerk.OrganizationInvitation, error) {
	return getClient().Revoke(ctx, params)
}

// ListAllFromInstance lists all the organization invitations from the current instance
func ListFromInstance(ctx context.Context, params *ListFromInstanceParams) (*clerk.OrganizationInvitationList, error) {
	return getClient().ListFromInstance(ctx, params)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}
