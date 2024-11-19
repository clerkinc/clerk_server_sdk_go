// Package commerce provides the Commerce API.
package commerce

import (
	"context"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Paths
const (
	path                = "/commerce"
	subscriptionsPath   = "/subscriptions"
	invoicesPath        = "/invoices"
	paymentAttemptsPath = "/payment_attempts"
	productsPath        = "/products"
	plansPath           = "/plans"
	integrationsPath    = "/integrations"
)

// Client is used to invoke the Commerce API.
type Client struct {
	Backend clerk.Backend
}

func NewClient(config *clerk.ClientConfig) *Client {
	return &Client{
		Backend: clerk.NewBackend(&config.BackendConfig),
	}
}

// Subscriptions

type ListSubscriptionsByInstanceIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// ListSubscriptionsByInstanceID returns a list of subscriptions for a given instance ID.
func (c *Client) ListSubscriptionsByInstanceID(ctx context.Context, params *ListSubscriptionsByInstanceIDParams) (*clerk.ListCommerceSubscriptionsResponse, error) {
	path, err := clerk.JoinPath(path, subscriptionsPath, params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.ListCommerceSubscriptionsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

type ListSubscriptionsByUserIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// ListSubscriptionsByUserID returns a list of subscriptions for a given user ID.
func (c *Client) ListSubscriptionsByUserID(ctx context.Context, params *ListSubscriptionsByUserIDParams) (*clerk.ListCommerceSubscriptionsResponse, error) {
	path, err := clerk.JoinPath(path, subscriptionsPath, "user", params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.ListCommerceSubscriptionsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

type GetSubscriptionByIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// GetSubscriptionByID retrieves a subscription by its ID.
func (c *Client) GetSubscriptionByID(ctx context.Context, params *GetSubscriptionByIDParams) (*clerk.CommerceSubscription, error) {
	path, err := clerk.JoinPath(path, subscriptionsPath, params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.CommerceSubscription{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// Invoices

type ListInvoicesByInstanceIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// ListInvoicesByInstanceID returns a list of invoices for a given instance ID.
func (c *Client) ListInvoicesByInstanceID(ctx context.Context, params *ListInvoicesByInstanceIDParams) (*clerk.ListCommerceInvoicesResponse, error) {
	path, err := clerk.JoinPath(path, invoicesPath, "instance", params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.ListCommerceInvoicesResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

type GetInvoiceByIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// GetInvoiceByID retrieves an invoice by its ID.
func (c *Client) GetInvoiceByID(ctx context.Context, params *GetInvoiceByIDParams) (*clerk.CommerceInvoice, error) {
	path, err := clerk.JoinPath(path, invoicesPath, params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.CommerceInvoice{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// Payment Attempts

type ListPaymentAttemptsByInstanceIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// ListPaymentAttemptsByInstanceID returns a list of payment attempts for a given instance ID.
func (c *Client) ListPaymentAttemptsByInstanceID(ctx context.Context, params *ListPaymentAttemptsByInstanceIDParams) (*clerk.ListCommercePaymentAttemptsResponse, error) {
	path, err := clerk.JoinPath(path, paymentAttemptsPath, "instance", params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.ListCommercePaymentAttemptsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

type GetPaymentAttemptByIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// GetPaymentAttemptByID retrieves a payment attempt by its ID.
func (c *Client) GetPaymentAttemptByID(ctx context.Context, params *GetPaymentAttemptByIDParams) (*clerk.CommercePaymentAttempt, error) {
	path, err := clerk.JoinPath(path, paymentAttemptsPath, params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.CommercePaymentAttempt{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// Products

type ListProductsByInstanceIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// ListProductsByInstanceID returns a list of products for a given instance ID.
func (c *Client) ListProductsByInstanceID(ctx context.Context, params *ListProductsByInstanceIDParams) (*clerk.ListCommerceProductsResponse, error) {
	path, err := clerk.JoinPath(path, productsPath, "instance", params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.ListCommerceProductsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// Plans

type GetPlanByIDParams struct {
	clerk.APIParams
	ID string `json:"-"`
}

// GetPlanByID retrieves a plan by its ID.
func (c *Client) GetPlanByID(ctx context.Context, params *GetPlanByIDParams) (*clerk.CommercePlan, error) {
	path, err := clerk.JoinPath(path, plansPath, params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, path)
	resource := &clerk.CommercePlan{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// Integrations

type CreateIntegrationParams struct {
	clerk.APIParams
	InstanceID string `json:"instance_id"`
	Email      string `json:"email"`
	Type       string `json:"type"`
}

// CreateIntegration creates a new integration for the specified instance.
func (c *Client) CreateIntegration(ctx context.Context, params *CreateIntegrationParams) (*clerk.CommerceIntegrationResponse, error) {
	path, err := clerk.JoinPath(path, integrationsPath)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, path)
	req.Body = params
	resource := &clerk.CommerceIntegrationResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
