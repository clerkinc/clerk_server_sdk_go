package integration

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../../cmd/gen/main.go

// Paths
const (
	rootPath = "/commerce"
	path     = "/integrations"
)

type CreateParams struct {
	clerk.APIParams
	Email           *string `json:"email,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty"`
}

type ListParams struct {
	clerk.APIParams
}

type Client struct {
	Backend clerk.Backend
}

func NewClient(config *clerk.ClientConfig) *Client {
	return &Client{
		Backend: clerk.NewBackend(&config.BackendConfig),
	}
}

func (c *Client) Create(ctx context.Context, params *CreateParams) (*clerk.CommerceIntegration, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceIntegration{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) List(ctx context.Context) (*clerk.CommerceIntegrationList, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceIntegrationList{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Get(ctx context.Context, id string) (*clerk.CommerceIntegration, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceIntegration{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

type UpdateParams struct {
	clerk.APIParams
	Status *string `json:"status,omitempty"`
}

func (c *Client) Update(ctx context.Context, id string, params *UpdateParams) (*clerk.CommerceIntegration, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceIntegration{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
