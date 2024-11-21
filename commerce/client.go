package commerce

import (
	"context"
	"time"

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

// Helper function for creating a *time.Time from a time.Time
func TimePtr(t time.Time) *time.Time {
	return &t
}

// --- Subscriptions ---
func (c *Client) CreateSubscription(ctx context.Context, params *clerk.CreateSubscriptionParams) (*clerk.CommerceSubscription, error) {
	return &clerk.CommerceSubscription{
		Customer: &clerk.CommerceCustomer{
			Entity: &struct {
				ID   *string `json:"id,omitempty"`
				Name *string `json:"name,omitempty"`
			}{
				ID:   clerk.String("customer_123"),
				Name: clerk.String("Alice Smith"),
			},
		},
		Plan: &clerk.CommercePlan{
			Name:        clerk.String("Basic Plan"),
			BaseAmount:  clerk.Int64(1000),
			IsRecurring: clerk.Bool(true),
		},
		Status: clerk.String("active"),
	}, nil
}

func (c *Client) ListSubscriptionsByInstanceID(ctx context.Context, params *clerk.ListSubscriptionsByInstanceIDParams) (*clerk.ListCommerceSubscriptionsResponse, error) {
	fakeSubscriptions := []clerk.CommerceSubscription{
		{
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
				Name:        clerk.String("Starter Plan"),
				BaseAmount:  clerk.Int64(1500),
				IsRecurring: clerk.Bool(true),
			},
			Status: clerk.String("active"),
		},
	}
	return &clerk.ListCommerceSubscriptionsResponse{
		PaginatedList: clerk.PaginatedList[clerk.CommerceSubscription]{
			Data:       &fakeSubscriptions,
			TotalCount: clerk.Int64(1),
		},
	}, nil
}

func (c *Client) ListSubscriptionsByUserID(ctx context.Context, params *clerk.ListSubscriptionsByUserIDParams) (*clerk.ListCommerceSubscriptionsResponse, error) {
	fakeSubscriptions := []clerk.CommerceSubscription{
		{
			Customer: &clerk.CommerceCustomer{
				Entity: &struct {
					ID   *string `json:"id,omitempty"`
					Name *string `json:"name,omitempty"`
				}{
					ID:   clerk.String("customer_2"),
					Name: clerk.String("Jane Doe"),
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
			TotalCount: clerk.Int64(1),
		},
	}, nil
}

func (c *Client) GetSubscriptionByID(ctx context.Context, params *clerk.GetSubscriptionByIDParams) (*clerk.CommerceSubscription, error) {
	return &clerk.CommerceSubscription{
		Customer: &clerk.CommerceCustomer{
			Entity: &struct {
				ID   *string `json:"id,omitempty"`
				Name *string `json:"name,omitempty"`
			}{
				ID:   clerk.String("customer_123"),
				Name: clerk.String("Alice Smith"),
			},
		},
		Plan: &clerk.CommercePlan{
			Name:        clerk.String("Premium Plan"),
			BaseAmount:  clerk.Int64(3000),
			IsRecurring: clerk.Bool(true),
		},
		Status: clerk.String("active"),
	}, nil
}

func (c *Client) UpdateSubscription(ctx context.Context, params *clerk.UpdateSubscriptionParams) (*clerk.CommerceSubscription, error) {
	return &clerk.CommerceSubscription{
		Customer: &clerk.CommerceCustomer{
			Entity: &struct {
				ID   *string `json:"id,omitempty"`
				Name *string `json:"name,omitempty"`
			}{
				ID:   clerk.String("customer_123"),
				Name: clerk.String("Alice Smith"),
			},
		},
		Plan: &clerk.CommercePlan{
			Name:        clerk.String("Updated Plan"),
			BaseAmount:  clerk.Int64(3500),
			IsRecurring: clerk.Bool(true),
		},
		Status: clerk.String("updated"),
	}, nil
}

// --- Invoices ---
func (c *Client) CreateInvoice(ctx context.Context, params *clerk.CreateInvoiceParams) (*clerk.CommerceInvoice, error) {
	return &clerk.CommerceInvoice{
		Subscription: &clerk.CommerceSubscription{
			Status: clerk.String("active"),
		},
		Amount: clerk.Int64(1000),
		Status: clerk.String("unpaid"),
		DueAt:  TimePtr(time.Now().Add(7 * 24 * time.Hour)),
	}, nil
}

func (c *Client) ListInvoicesByInstanceID(ctx context.Context, params *clerk.ListInvoicesByInstanceIDParams) (*clerk.ListCommerceInvoicesResponse, error) {
	fakeInvoices := []clerk.CommerceInvoice{
		{
			Amount: clerk.Int64(5000),
			Status: clerk.String("paid"),
			DueAt:  TimePtr(time.Now().Add(14 * 24 * time.Hour)),
		},
	}
	return &clerk.ListCommerceInvoicesResponse{
		PaginatedList: clerk.PaginatedList[clerk.CommerceInvoice]{
			Data:       &fakeInvoices,
			TotalCount: clerk.Int64(1),
		},
	}, nil
}

func (c *Client) GetInvoiceByID(ctx context.Context, params *clerk.GetInvoiceByIDParams) (*clerk.CommerceInvoice, error) {
	return &clerk.CommerceInvoice{
		Amount: clerk.Int64(7000),
		Status: clerk.String("due"),
		DueAt:  TimePtr(time.Now().Add(5 * 24 * time.Hour)),
	}, nil
}

func (c *Client) UpdateInvoice(ctx context.Context, params *clerk.UpdateInvoiceParams) (*clerk.CommerceInvoice, error) {
	return &clerk.CommerceInvoice{
		Amount: clerk.Int64(8000),
		Status: clerk.String("overdue"),
		DueAt:  TimePtr(time.Now().Add(-2 * 24 * time.Hour)),
	}, nil
}

// --- Payment Attempts ---
func (c *Client) CreatePaymentAttempt(ctx context.Context, params *clerk.CreatePaymentAttemptParams) (*clerk.CommercePaymentAttempt, error) {
	return &clerk.CommercePaymentAttempt{
		Amount: clerk.Int64(1000),
		Status: clerk.String("failed"),
	}, nil
}

func (c *Client) ListPaymentAttemptsByInstanceID(ctx context.Context, params *clerk.ListPaymentAttemptsByInstanceIDParams) (*clerk.ListCommercePaymentAttemptsResponse, error) {
	fakePaymentAttempts := []clerk.CommercePaymentAttempt{
		{
			Amount: clerk.Int64(5000),
			Status: clerk.String("success"),
		},
	}
	return &clerk.ListCommercePaymentAttemptsResponse{
		PaginatedList: clerk.PaginatedList[clerk.CommercePaymentAttempt]{
			Data:       &fakePaymentAttempts,
			TotalCount: clerk.Int64(1),
		},
	}, nil
}

func (c *Client) GetPaymentAttemptByID(ctx context.Context, params *clerk.GetPaymentAttemptByIDParams) (*clerk.CommercePaymentAttempt, error) {
	return &clerk.CommercePaymentAttempt{
		Amount: clerk.Int64(3000),
		Status: clerk.String("pending"),
	}, nil
}

func (c *Client) UpdatePaymentAttempt(ctx context.Context, params *clerk.UpdatePaymentAttemptParams) (*clerk.CommercePaymentAttempt, error) {
	return &clerk.CommercePaymentAttempt{
		Amount: clerk.Int64(4000),
		Status: clerk.String("refunded"),
	}, nil
}

// --- Products ---
func (c *Client) CreateProduct(ctx context.Context, params *clerk.CreateProductParams) (*clerk.CommerceProduct, error) {
	return &clerk.CommerceProduct{
		Name:            *clerk.String("New Product"),
		Currency:        *clerk.String("USD"),
		SubscriberType:  *clerk.String("individual"),
		OwnerEntityType: *clerk.String("business"),
	}, nil
}

func (c *Client) ListProductsByInstanceID(ctx context.Context, params *clerk.ListProductsByInstanceIDParams) (*clerk.ListCommerceProductsResponse, error) {
	fakeProducts := []clerk.CommerceProduct{
		{
			Name:            *clerk.String("Product 1"),
			Currency:        *clerk.String("USD"),
			SubscriberType:  *clerk.String("organization"),
			OwnerEntityType: *clerk.String("enterprise"),
		},
	}
	return &clerk.ListCommerceProductsResponse{
		PaginatedList: clerk.PaginatedList[clerk.CommerceProduct]{
			Data:       &fakeProducts,
			TotalCount: clerk.Int64(1),
		},
	}, nil
}

func (c *Client) GetProductByID(ctx context.Context, params *clerk.GetProductByIDParams) (*clerk.CommerceProduct, error) {
	return &clerk.CommerceProduct{
		Name:            *clerk.String("Specific Product"),
		Currency:        *clerk.String("EUR"),
		SubscriberType:  *clerk.String("group"),
		OwnerEntityType: *clerk.String("nonprofit"),
	}, nil
}

func (c *Client) UpdateProduct(ctx context.Context, params *clerk.UpdateProductParams) (*clerk.CommerceProduct, error) {
	return &clerk.CommerceProduct{
		Name:            *clerk.String("Updated Product"),
		Currency:        *clerk.String("GBP"),
		SubscriberType:  *clerk.String("individual"),
		OwnerEntityType: *clerk.String("corporate"),
	}, nil
}

// --- Plans ---
func (c *Client) CreatePlan(ctx context.Context, params *clerk.CreatePlanParams) (*clerk.CommercePlan, error) {
	return &clerk.CommercePlan{
		Name:        clerk.String("Basic Plan"),
		BaseAmount:  clerk.Int64(1200),
		IsRecurring: clerk.Bool(true),
	}, nil
}

func (c *Client) ListPlansByInstanceID(ctx context.Context, params *clerk.ListPlansByInstanceIDParams) (*clerk.ListCommerceProductsResponse, error) {
	fakeProducts := []clerk.CommerceProduct{
		{
			Name:            *clerk.String("Product A"),
			Slug:            *clerk.String("product-a"),
			Currency:        *clerk.String("USD"),
			SubscriberType:  "individual",
			OwnerEntityType: *clerk.String("business"),
		},
		{
			Name:            *clerk.String("Product B"),
			Slug:            *clerk.String("product-b"),
			Currency:        *clerk.String("EUR"),
			SubscriberType:  "organization",
			OwnerEntityType: *clerk.String("enterprise"),
		},
	}
	return &clerk.ListCommerceProductsResponse{
		PaginatedList: clerk.PaginatedList[clerk.CommerceProduct]{
			Data:       &fakeProducts,                         // Ensure this is a slice of CommerceProduct
			TotalCount: clerk.Int64(int64(len(fakeProducts))), // Adjust the count based on the slice
		},
	}, nil
}

func (c *Client) GetPlanByID(ctx context.Context, params *clerk.GetPlanByIDParams) (*clerk.CommercePlan, error) {
	return &clerk.CommercePlan{
		Name:        clerk.String("Exclusive Plan"),
		BaseAmount:  clerk.Int64(5000),
		IsRecurring: clerk.Bool(false),
	}, nil
}

func (c *Client) UpdatePlan(ctx context.Context, params *clerk.UpdatePlanParams) (*clerk.CommercePlan, error) {
	return &clerk.CommercePlan{
		Name:        clerk.String("Updated Plan"),
		BaseAmount:  clerk.Int64(5500),
		IsRecurring: clerk.Bool(true),
	}, nil
}

// --- Integrations ---
func (c *Client) CreateIntegration(ctx context.Context, params *clerk.CreateIntegrationParams) (*clerk.CommerceIntegrationResponse, error) {
	return &clerk.CommerceIntegrationResponse{
		URL: clerk.String("https://fake-integration-url.com"),
	}, nil
}

func (c *Client) ListIntegrationsByInstanceID(ctx context.Context, params *clerk.ListIntegrationsByInstanceIDParams) (*clerk.ListCommerceIntegrationsResponse, error) {
	fakeIntegrations := []clerk.CommerceIntegration{
		{
			IntegrationID:   clerk.String("int_123"),
			IntegrationType: clerk.String("type_a"),
			Status:          clerk.String("active"),
		},
	}
	return &clerk.ListCommerceIntegrationsResponse{
		PaginatedList: clerk.PaginatedList[clerk.CommerceIntegration]{
			Data:       &fakeIntegrations,
			TotalCount: clerk.Int64(1),
		},
	}, nil
}

func (c *Client) GetIntegration(ctx context.Context, params *clerk.GetIntegrationParams) (*clerk.CommerceIntegrationResponse, error) {
	return &clerk.CommerceIntegrationResponse{
		URL: clerk.String("https://fake-integration-detail.com"),
	}, nil
}

func (c *Client) UpdateIntegration(ctx context.Context, params *clerk.UpdateIntegrationParams) (*clerk.CommerceIntegrationResponse, error) {
	return &clerk.CommerceIntegrationResponse{
		URL: clerk.String("https://fake-integration-updated.com"),
	}, nil
}
