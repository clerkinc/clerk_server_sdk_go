// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package waitlistentry

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

// List returns all waitlist entries.
func List(ctx context.Context, params *ListParams) (*clerk.WaitlistEntriesList, error) {
	return getClient().List(ctx, params)
}

// Create adds a new waitlist entry.
func Create(ctx context.Context, params *CreateParams) (*clerk.WaitlistEntry, error) {
	return getClient().Create(ctx, params)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}