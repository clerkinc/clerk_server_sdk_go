package paymentsources

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../../cmd/gen/main.go
const (
	rootPath = "/commerce"
	path     = "/payment_sources"
)

type CreateParams struct {
	clerk.APIParams
	OwnerID   string `json:"owner_id"`
	GatewayID string `json:"gateway_id"`
	Processor string `json:"processor"`
}

type UpdateParams struct {
	clerk.APIParams
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

func (c *Client) Create(ctx context.Context, params *CreateParams) (*clerk.CommercePaymentSource, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePaymentSource{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) List(ctx context.Context, params *ListParams) (*clerk.CommercePaymentSourceList, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePaymentSourceList{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Get(ctx context.Context, id string) (*clerk.CommercePaymentSource, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePaymentSource{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Update(ctx context.Context, id string, params *UpdateParams) (*clerk.CommercePaymentSource, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePaymentSource{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
