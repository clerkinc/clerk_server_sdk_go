// Package user provides the Users API.
package commerce

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../cmd/gen/main.go

const path = "/commerce"
const subscriptionsPath = "/subscriptions"

// Client is used to invoke the Users API.
type Client struct {
	Backend clerk.Backend
}

func NewClient(config *clerk.ClientConfig) *Client {
	return &Client{
		Backend: clerk.NewBackend(&config.BackendConfig),
	}
}

type ListSubscriptionsByInstanceIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// ListSubscriptionsByInstanceID returns a list of subscriptions for a given instance ID.
func (c *Client) ListSubscriptionsByInstanceID(ctx context.Context, params *ListSubscriptionsByInstanceIDParams) (*clerk.SubscriptionList, error) {
	path, err := clerk.JoinPath(path, subscriptionsPath, params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.SubscriptionList{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
