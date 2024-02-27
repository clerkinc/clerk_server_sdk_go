// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package samlconnection

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Create creates a new SAML Connection.
func Create(ctx context.Context, params *CreateParams) (*clerk.SAMLConnection, error) {
	return getClient().Create(ctx, params)
}

// Get returns details about a SAML Connection.
func Get(ctx context.Context, id string) (*clerk.SAMLConnection, error) {
	return getClient().Get(ctx, id)
}

// Update updates the SAML Connection specified by id.
func Update(ctx context.Context, id string, params *UpdateParams) (*clerk.SAMLConnection, error) {
	return getClient().Update(ctx, id, params)
}

// Delete deletes a SAML Connection.
func Delete(ctx context.Context, id string) (*clerk.DeletedResource, error) {
	return getClient().Delete(ctx, id)
}

// List returns a list of SAML Connections.
func List(ctx context.Context, params *ListParams) (*clerk.SAMLConnectionList, error) {
	return getClient().List(ctx, params)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}
