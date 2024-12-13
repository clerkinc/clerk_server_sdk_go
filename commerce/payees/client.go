
package payee

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../../cmd/gen/main.go
const (
	rootPath = "/commerce"
	path     = "/payees"
)

type CreateParams struct {
	clerk.APIParams
	SubscriptionID *string `json:"subscription_id,omitempty"`
	Amount         *int64  `json:"amount,omitempty"`
	DueAt          *string `json:"due_at,omitempty"`
}

type UpdateParams struct {
	clerk.APIParams
	Status *string `json:"status,omitempty"`
}

type ListParams struct {
	clerk.APIParams
	InstanceID *string `json:"-"`
}

type Client struct {
	Backend clerk.Backend
}

func NewClient(config *clerk.ClientConfig) *Client {
	return &Client{
		Backend: clerk.NewBackend(&config.BackendConfig),
	}
}

func (c *Client) Create(ctx context.Context, params *CreateParams) (*clerk.CommercePayee, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePayee{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) List(ctx context.Context, params *ListParams) (*clerk.CommercePayeeList, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePayeeList{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Get(ctx context.Context, id string) (*clerk.CommercePayee, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePayee{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Update(ctx context.Context, id string, params *UpdateParams) (*clerk.CommercePayee, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePayee{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
