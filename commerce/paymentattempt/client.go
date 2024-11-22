package paymentattempt

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../../cmd/gen/main.go
const (
	rootPath = "/commerce"
	path     = "/payment_attempts"
)

type Client struct {
	Backend clerk.Backend
}

func NewClient(config *clerk.ClientConfig) *Client {
	return &Client{
		Backend: clerk.NewBackend(&config.BackendConfig),
	}
}

type CreateParams struct {
	clerk.APIParams
	InvoiceID *string `json:"invoice_id,omitempty"`
}

func (c *Client) Create(ctx context.Context, params *CreateParams) (*clerk.CommercePaymentAttempt, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePaymentAttempt{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

type ListParams struct {
	clerk.APIParams
}

func (c *Client) List(ctx context.Context, params *ListParams) (*clerk.ListCommercePaymentAttemptsResponse, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommercePaymentAttemptsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Get(ctx context.Context, id string) (*clerk.CommercePaymentAttempt, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePaymentAttempt{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

type UpdateParams struct {
	clerk.APIParams
	Status *string `json:"status,omitempty"`
}

func (c *Client) Update(ctx context.Context, id string, params *UpdateParams) (*clerk.CommercePaymentAttempt, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePaymentAttempt{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
