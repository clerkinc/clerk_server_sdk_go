package subscriptions

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../../cmd/gen/main.go

const (
	rootPath = "/commerce"
	path     = "/subscriptions"
)

type Client struct {
	Backend clerk.Backend
}

func NewClient(config *clerk.ClientConfig) *Client {
	return &Client{
		Backend: clerk.NewBackend(&config.BackendConfig),
	}
}

func (c *Client) Create(ctx context.Context, params *clerk.CreateSubscriptionParams) (*clerk.CommerceSubscription, error) {
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceSubscription{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) List(ctx context.Context, includes []string) (*clerk.ListCommerceSubscriptionsResponse, error) {
	// Build the base path
	reqPath, err := clerk.JoinPath(rootPath, path)
	if err != nil {
		return nil, err
	}

	if len(includes) > 0 {
		query := url.Values{}
		query.Set("include", strings.Join(includes, ","))
		reqPath += "?" + query.Encode()
	}

	// Create the API request
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceSubscriptionsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListInvoices(ctx context.Context, subscriptionID string) (*clerk.ListCommerceSubscriptionsResponse, error) {
	// Build the base path
	reqPath, err := clerk.JoinPath(rootPath, path, subscriptionID, "invoices")
	if err != nil {
		return nil, err
	}

	// Create the API request
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceSubscriptionsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListByUserID(ctx context.Context, params *clerk.ListSubscriptionsByUserIDParams) (*clerk.ListCommerceSubscriptionsResponse, error) {
	reqPath, err := clerk.JoinPath(rootPath, "subscribers", params.SubscriberType, params.ID, "subscriptions")
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceSubscriptionsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Get(ctx context.Context, id string) (*clerk.CommerceSubscription, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceSubscription{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) Update(ctx context.Context, id string, params *clerk.UpdateSubscriptionParams) (*clerk.CommerceSubscription, error) {
	reqPath, err := clerk.JoinPath(rootPath, path, id)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceSubscription{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
