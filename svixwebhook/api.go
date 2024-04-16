// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package svixwebhook

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Create creates a Svix app.
func Create(ctx context.Context) (*clerk.SvixWebhook, error) {
	return getClient().Create(ctx)
}

// Delete deletes the Svix app.
func Delete(ctx context.Context) (*clerk.SvixWebhook, error) {
	return getClient().Delete(ctx)
}

// RefreshURL generates a new URL for accessing Svix's dashboard.
func RefreshURL(ctx context.Context) (*clerk.SvixWebhook, error) {
	return getClient().RefreshURL(ctx)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}