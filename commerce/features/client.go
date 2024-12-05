package features

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../../cmd/gen/main.go
const (
	rootPath = "/commerce"
	path     = "/features"
)

type Client struct {
	Backend clerk.Backend
}

func NewClient(config *clerk.ClientConfig) *Client {
	return &Client{
		Backend: clerk.NewBackend(&config.BackendConfig),
	}
}

func (c *Client) Create(ctx context.Context, params *clerk.CreateFeatureParams) (*clerk.CommerceFeature, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceFeature{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) List(ctx context.Context, params *clerk.ListFeaturesByInstanceIDParams) (*clerk.CommerceFeatureList, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceFeatureList{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListByPlanID(ctx context.Context, params *clerk.ListFeaturesByPlanIDParams) (*clerk.CommerceFeatureList, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceFeatureList{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Get(ctx context.Context, id string) (*clerk.CommerceFeature, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceFeature{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Delete(ctx context.Context, id string) error {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return err
	}
	req := clerk.NewAPIRequest(http.MethodDelete, reqPath)
	err = c.Backend.Call(ctx, req, nil)
	return err
}

func (c *Client) Update(ctx context.Context, id string, params *clerk.UpdateFeatureParams) (*clerk.CommerceFeature, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceFeature{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
