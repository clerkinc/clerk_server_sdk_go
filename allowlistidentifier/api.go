// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package allowlistidentifier

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v3"
)

// Create adds a new identifier to the allowlist.
func Create(ctx context.Context, params *CreateParams) (*clerk.AllowlistIdentifier, error) {
	return getClient().Create(ctx, params)
}

// Delete removes an identifier from the allowlist.
func Delete(ctx context.Context, id string) (*clerk.DeletedResource, error) {
	return getClient().Delete(ctx, id)
}

// List returns all the identifiers in the allowlist.
func List(ctx context.Context, params *ListParams) (*clerk.AllowlistIdentifierList, error) {
	return getClient().List(ctx, params)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}
