// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package invitation

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v3"
)

// List returns all invitations.
func List(ctx context.Context, params *ListParams) (*clerk.InvitationList, error) {
	return getClient().List(ctx, params)
}

// Create adds a new identifier to the allowlist.
func Create(ctx context.Context, params *CreateParams) (*clerk.Invitation, error) {
	return getClient().Create(ctx, params)
}

// BulkCreate creates multiple invitations.
func BulkCreate(ctx context.Context, params *BulkCreateParams) (*clerk.Invitations, error) {
	return getClient().BulkCreate(ctx, params)
}

// Revoke revokes a pending invitation.
func Revoke(ctx context.Context, id string) (*clerk.Invitation, error) {
	return getClient().Revoke(ctx, id)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}
