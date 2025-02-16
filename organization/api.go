// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package organization

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Create creates a new organization.
func Create(ctx context.Context, params *CreateParams) (*clerk.Organization, error) {
	return getClient().Create(ctx, params)
}

// Get retrieves details for an organization.
// The organization can be fetched by either the ID or its slug.
func Get(ctx context.Context, idOrSlug string) (*clerk.Organization, error) {
	return getClient().Get(ctx, idOrSlug)
}

// GetWithParams retrieves details for an organization.
// The organization can be fetched by either the ID or its slug.
func GetWithParams(ctx context.Context, idOrSlug string, params *GetParams) (*clerk.Organization, error) {
	return getClient().GetWithParams(ctx, idOrSlug, params)
}

// Update updates an organization.
func Update(ctx context.Context, id string, params *UpdateParams) (*clerk.Organization, error) {
	return getClient().Update(ctx, id, params)
}

// UpdateMetadata updates the organization's metadata by merging the
// provided values with the existing ones.
func UpdateMetadata(ctx context.Context, id string, params *UpdateMetadataParams) (*clerk.Organization, error) {
	return getClient().UpdateMetadata(ctx, id, params)
}

// Delete deletes an organization.
func Delete(ctx context.Context, id string) (*clerk.DeletedResource, error) {
	return getClient().Delete(ctx, id)
}

// UpdateLogo sets or replaces the organization's logo.
func UpdateLogo(ctx context.Context, id string, params *UpdateLogoParams) (*clerk.Organization, error) {
	return getClient().UpdateLogo(ctx, id, params)
}

// DeleteLogo removes the organization's logo.
func DeleteLogo(ctx context.Context, id string) (*clerk.Organization, error) {
	return getClient().DeleteLogo(ctx, id)
}

// List returns a list of organizations.
func List(ctx context.Context, params *ListParams) (*clerk.OrganizationList, error) {
	return getClient().List(ctx, params)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}
