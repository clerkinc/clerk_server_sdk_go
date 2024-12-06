package plans

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../../cmd/gen/main.go
const (
	rootPath = "/commerce"
	path     = "/plans"
)

type Client struct {
	Backend clerk.Backend
}

func NewClient(config *clerk.ClientConfig) *Client {
	return &Client{
		Backend: clerk.NewBackend(&config.BackendConfig),
	}
}

func (c *Client) Create(ctx context.Context, params *clerk.CreatePlanParams) (*clerk.CommercePlan, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePlan{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) List(ctx context.Context, params *clerk.ListPlansByInstanceIDParams) (*clerk.CommercePlanList, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePlanList{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Get(ctx context.Context, id string) (*clerk.CommercePlan, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePlan{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Update(ctx context.Context, id string, params *clerk.UpdatePlanParams) (*clerk.CommercePlan, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePlan{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
