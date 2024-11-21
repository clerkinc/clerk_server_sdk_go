package commerce

import (
	"context"
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

//go:generate go run ../cmd/gen/main.go

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

// --- Subscriptions ---
func (c *Client) CreateSubscription(ctx context.Context, params *clerk.CreateSubscriptionParams) (*clerk.CommerceSubscription, error) {
	reqPath, err := clerk.JoinPath(path, subscriptionsPath)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceSubscription{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListSubscriptionsByInstanceID(ctx context.Context, params *clerk.ListSubscriptionsByInstanceIDParams) (*clerk.ListCommerceSubscriptionsResponse, error) {
	reqPath, err := clerk.JoinPath(path, subscriptionsPath, "instance", *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceSubscriptionsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	log.Default().Println("resource", resource)
	// Fake data for testing
	fakeSubscriptions := []clerk.CommerceSubscription{
		{
			APIResource: clerk.APIResource{}, // Fill in with appropriate resource metadata if needed
			Customer: &clerk.CommerceCustomer{
				Entity: &struct {
					ID   *string `json:"id,omitempty"`
					Name *string `json:"name,omitempty"`
				}{
					ID:   clerk.String("customer_1"),
					Name: clerk.String("John Doe"),
				},
			},
			Plan: &clerk.CommercePlan{
				Name:        clerk.String("Basic Plan"),
				BaseAmount:  clerk.Int64(1000),
				IsRecurring: clerk.Bool(true),
			},
			Status: clerk.String("active"),
		},
		{
			APIResource: clerk.APIResource{},
			Customer: &clerk.CommerceCustomer{
				Entity: &struct {
					ID   *string `json:"id,omitempty"`
					Name *string `json:"name,omitempty"`
				}{
					ID:   clerk.String("customer_2"),
					Name: clerk.String("Jane Smith"),
				},
			},
			Plan: &clerk.CommercePlan{
				Name:        clerk.String("Pro Plan"),
				BaseAmount:  clerk.Int64(2000),
				IsRecurring: clerk.Bool(true),
			},
			Status: clerk.String("inactive"),
		},
	}

	return &clerk.ListCommerceSubscriptionsResponse{
		PaginatedList: clerk.PaginatedList[clerk.CommerceSubscription]{
			Data:       &fakeSubscriptions,
			TotalCount: clerk.Int64(int64(len(fakeSubscriptions))),
		},
	}, nil
	// return resource, err
}

func (c *Client) ListSubscriptionsByUserID(ctx context.Context, params *clerk.ListSubscriptionsByUserIDParams) (*clerk.ListCommerceSubscriptionsResponse, error) {
	reqPath, err := clerk.JoinPath(path, subscriptionsPath, "user", *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceSubscriptionsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) GetSubscriptionByID(ctx context.Context, params *clerk.GetSubscriptionByIDParams) (*clerk.CommerceSubscription, error) {
	reqPath, err := clerk.JoinPath(path, subscriptionsPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceSubscription{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) UpdateSubscription(ctx context.Context, params *clerk.UpdateSubscriptionParams) (*clerk.CommerceSubscription, error) {
	reqPath, err := clerk.JoinPath(path, subscriptionsPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceSubscription{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// --- Invoices ---
func (c *Client) CreateInvoice(ctx context.Context, params *clerk.CreateInvoiceParams) (*clerk.CommerceInvoice, error) {
	reqPath, err := clerk.JoinPath(path, invoicesPath)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceInvoice{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListInvoicesByInstanceID(ctx context.Context, params *clerk.ListInvoicesByInstanceIDParams) (*clerk.ListCommerceInvoicesResponse, error) {
	reqPath, err := clerk.JoinPath(path, invoicesPath, "instance", *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceInvoicesResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) GetInvoiceByID(ctx context.Context, params *clerk.GetInvoiceByIDParams) (*clerk.CommerceInvoice, error) {
	reqPath, err := clerk.JoinPath(path, invoicesPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceInvoice{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) UpdateInvoice(ctx context.Context, params *clerk.UpdateInvoiceParams) (*clerk.CommerceInvoice, error) {
	reqPath, err := clerk.JoinPath(path, invoicesPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceInvoice{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// --- Payment Attempts ---
func (c *Client) CreatePaymentAttempt(ctx context.Context, params *clerk.CreatePaymentAttemptParams) (*clerk.CommercePaymentAttempt, error) {
	reqPath, err := clerk.JoinPath(path, paymentAttemptsPath)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePaymentAttempt{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListPaymentAttemptsByInstanceID(ctx context.Context, params *clerk.ListPaymentAttemptsByInstanceIDParams) (*clerk.ListCommercePaymentAttemptsResponse, error) {
	reqPath, err := clerk.JoinPath(path, paymentAttemptsPath, "instance", *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommercePaymentAttemptsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) GetPaymentAttemptByID(ctx context.Context, params *clerk.GetPaymentAttemptByIDParams) (*clerk.CommercePaymentAttempt, error) {
	reqPath, err := clerk.JoinPath(path, paymentAttemptsPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePaymentAttempt{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) UpdatePaymentAttempt(ctx context.Context, params *clerk.UpdatePaymentAttemptParams) (*clerk.CommercePaymentAttempt, error) {
	reqPath, err := clerk.JoinPath(path, paymentAttemptsPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePaymentAttempt{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// --- Products ---
func (c *Client) CreateProduct(ctx context.Context, params *clerk.CreateProductParams) (*clerk.CommerceProduct, error) {
	reqPath, err := clerk.JoinPath(path, productsPath)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceProduct{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListProductsByInstanceID(ctx context.Context, params *clerk.ListProductsByInstanceIDParams) (*clerk.ListCommerceProductsResponse, error) {
	reqPath, err := clerk.JoinPath(path, productsPath, "instance", *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceProductsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) GetProductByID(ctx context.Context, params *clerk.GetProductByIDParams) (*clerk.CommerceProduct, error) {
	reqPath, err := clerk.JoinPath(path, productsPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceProduct{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) UpdateProduct(ctx context.Context, params *clerk.UpdateProductParams) (*clerk.CommerceProduct, error) {
	reqPath, err := clerk.JoinPath(path, productsPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceProduct{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// --- Plans ---
func (c *Client) CreatePlan(ctx context.Context, params *clerk.CreatePlanParams) (*clerk.CommercePlan, error) {
	reqPath, err := clerk.JoinPath(path, plansPath)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePlan{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListPlansByInstanceID(ctx context.Context, params *clerk.ListPlansByInstanceIDParams) (*clerk.ListCommerceProductsResponse, error) {
	reqPath, err := clerk.JoinPath(path, plansPath, "instance", *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceProductsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) GetPlanByID(ctx context.Context, params *clerk.GetPlanByIDParams) (*clerk.CommercePlan, error) {
	reqPath, err := clerk.JoinPath(path, plansPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommercePlan{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) UpdatePlan(ctx context.Context, params *clerk.UpdatePlanParams) (*clerk.CommercePlan, error) {
	reqPath, err := clerk.JoinPath(path, plansPath, *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommercePlan{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

// --- Integrations ---
func (c *Client) CreateIntegration(ctx context.Context, params *clerk.CreateIntegrationParams) (*clerk.CommerceIntegrationResponse, error) {
	reqPath, err := clerk.JoinPath(path, integrationsPath)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPost, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceIntegrationResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) ListIntegrationsByInstanceID(ctx context.Context, params *clerk.ListIntegrationsByInstanceIDParams) (*clerk.ListCommerceIntegrationsResponse, error) {
	reqPath, err := clerk.JoinPath(path, integrationsPath, "instance", *params.ID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.ListCommerceIntegrationsResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) GetIntegration(ctx context.Context, params *clerk.GetIntegrationParams) (*clerk.CommerceIntegrationResponse, error) {
	reqPath, err := clerk.JoinPath(path, integrationsPath, *params.IntegrationID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodGet, reqPath)
	resource := &clerk.CommerceIntegrationResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}

func (c *Client) UpdateIntegration(ctx context.Context, params *clerk.UpdateIntegrationParams) (*clerk.CommerceIntegrationResponse, error) {
	reqPath, err := clerk.JoinPath(path, integrationsPath, *params.CommerceIntegrationID)
	if err != nil {
		return nil, err
	}
	req := clerk.NewAPIRequest(http.MethodPut, reqPath)
	req.SetParams(params)
	resource := &clerk.CommerceIntegrationResponse{}
	err = c.Backend.Call(ctx, req, resource)
	return resource, err
}
