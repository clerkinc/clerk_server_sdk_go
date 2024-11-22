package clerk

import (
	"time"
)

// --- Product Types ---

type CreateProductParams struct {
	APIParams
	InstanceID      string   `json:"instance_id"`
	Name            string   `json:"name"`
	Slug            string   `json:"slug"`
	Currency        string   `json:"currency"`
	SubscriberType  []string `json:"subscriber_type"`
	OwnerEntityType string   `json:"owner_entity_type"`
}

type UpdateProductParams struct {
	APIParams
	ID              string    `json:"id"`
	Name            *string   `json:"name,omitempty"`
	Slug            *string   `json:"slug,omitempty"`
	Currency        *string   `json:"currency,omitempty"`
	SubscriberType  *[]string `json:"subscriber_type,omitempty"`
	OwnerEntityType *string   `json:"owner_entity_type,omitempty"`
}

type GetProductByIDParams struct {
	APIParams
	ID string `json:"id"`
}

type CommerceProduct struct {
	APIResource
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Currency        string    `json:"currency"`
	SubscriberType  []string  `json:"subscriber_type"`
	OwnerEntityType string    `json:"owner_entity_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CommerceProductWithPlans struct {
	CommerceProduct
	Plans []CommercePlan `json:"plans"`
}

type ListProductsByInstanceIDParams struct {
	APIParams
	ID string `json:"id"`
}

type ListCommerceProductsResponse struct {
	APIResource
	PaginatedList[CommerceProduct]
}

// --- Plan Types ---

type CreatePlanParams struct {
	APIParams
	Name        string `json:"name"`
	ProductID   string `json:"product_id"`
	BaseAmount  int64  `json:"base_amount"`
	IsRecurring bool   `json:"is_recurring"`
}

type UpdatePlanParams struct {
	APIParams
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetPlanByIDParams struct {
	APIParams
	ID string `json:"id"`
}

type CommercePlan struct {
	APIResource
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Product         CommerceProduct `json:"product"`
	BaseAmount      int64           `json:"base_amount"`
	IsRecurring     bool            `json:"is_recurring"`
	IsProrated      bool            `json:"is_prorated"`
	Period          string          `json:"period"` // e.g., "month" or "year"
	Interval        int             `json:"interval"`
	BillingCycles   *int            `json:"billing_cycles,omitempty"`
	SubscriberCount int64           `json:"subscriber_count"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type ListPlansByInstanceIDParams struct {
	APIParams
	ID string `json:"id"`
}

// --- Integration Types ---

type CreateIntegrationParams struct {
	APIParams
	InstanceID string `json:"instance_id"`
	Email      string `json:"email"`
	Type       string `json:"type"`
}

type UpdateIntegrationParams struct {
	APIParams
	CommerceIntegrationID string `json:"id"`
	Status                string `json:"status"`
}

type GetIntegrationParams struct {
	APIParams
	IntegrationID string `json:"id"`
}

type CommerceIntegration struct {
	APIResource
	IntegrationID   string `json:"integration_id"`
	IntegrationType string `json:"integration_type"`
	Status          string `json:"status"`
}

type CommerceIntegrationResponse struct {
	APIResource
	URL string `json:"url"`
}

type ListCommerceIntegrationsResponse struct {
	APIResource
	PaginatedList[CommerceIntegration]
}

type ListIntegrationsByInstanceIDParams struct {
	APIParams
	ID string `json:"id"`
}

// --- Subscription Types ---

type CreateSubscriptionParams struct {
	APIParams
	CustomerID string `json:"customer_id"`
	PlanID     string `json:"plan_id"`
	Status     string `json:"status"`
}

type UpdateSubscriptionParams struct {
	APIParams
	ID     string  `json:"id"`
	Status *string `json:"status,omitempty"`
}

type GetSubscriptionByIDParams struct {
	APIParams
	ID string `json:"id"`
}

type ListSubscriptionsByInstanceIDParams struct {
	APIParams
	ID string `json:"id"`
}

type ListSubscriptionsByUserIDParams struct {
	APIParams
	ID string `json:"id"`
}

type CommerceSubscription struct {
	APIResource
	ID          string           `json:"id"`
	AppID       string           `json:"app_id"`
	Customer    CommerceCustomer `json:"customer"`
	Plan        CommercePlan     `json:"plan"`
	Status      string           `json:"status"`
	LastInvoice *CommerceInvoice `json:"last_invoice,omitempty"`
	NextInvoice *CommerceInvoice `json:"next_invoice,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type ListCommerceSubscriptionsResponse struct {
	APIResource
	PaginatedList[CommerceSubscription]
}

// --- Subscription Types ---

type GetSubscriptionParams struct {
	APIParams
	ID string `json:"id"`
}

type ListSubscribersParams struct {
	APIParams
	InstanceID string `json:"instance_id"`
}

type CommerceSubscriber struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ListCommerceSubscribersResponse struct {
	APIResource
	PaginatedList[CommerceSubscriber]
}

// Supporting structs for emails
type ClerkEmail struct {
	Address string `json:"address"`
}

// --- Invoice Types ---

type CreateInvoiceParams struct {
	APIParams
	SubscriptionID string `json:"subscription_id"`
	Amount         int64  `json:"amount"`
	DueAt          string `json:"due_at"`
}

type UpdateInvoiceParams struct {
	APIParams
	ID     string  `json:"id"`
	Status *string `json:"status,omitempty"`
}

type GetInvoiceByIDParams struct {
	APIParams
	ID string `json:"id"`
}

type ListInvoicesByInstanceIDParams struct {
	APIParams
	ID string `json:"id"`
}

type CommerceInvoice struct {
	APIResource
	ID                       string                `json:"id"`
	Subscription             *CommerceSubscription `json:"subscription,omitempty"`
	Amount                   int64                 `json:"amount"`
	Status                   string                `json:"status"`
	DueAt                    *time.Time            `json:"due_at,omitempty"`
	FinalizingPaymentAttempt string                `json:"finalizing_payment_attempt_id,omitempty"`
}

type ListCommerceInvoicesResponse struct {
	APIResource
	PaginatedList[CommerceInvoice]
}

// --- Payment Attempt Types ---

type CreatePaymentAttemptParams struct {
	APIParams
	InvoiceID string `json:"invoice_id"`
	Amount    int64  `json:"amount"`
	Status    string `json:"status"`
}

type UpdatePaymentAttemptParams struct {
	APIParams
	ID     string  `json:"id"`
	Status *string `json:"status,omitempty"`
}

type GetPaymentAttemptByIDParams struct {
	APIParams
	ID string `json:"id"`
}

type ListPaymentAttemptsByInstanceIDParams struct {
	APIParams
	ID string `json:"id"`
}

type CommercePaymentAttempt struct {
	APIResource
	ID        string          `json:"id"`
	Invoice   CommerceInvoice `json:"invoice"`
	Amount    int64           `json:"amount"`
	Status    string          `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
}

type ListCommercePaymentAttemptsResponse struct {
	APIResource
	PaginatedList[CommercePaymentAttempt]
}

// --- Customer Types ---

type CommerceCustomer struct {
	ID     string `json:"id"`
	AppID  string `json:"app_id"`
	Entity *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"entity"`
}

// --- Pagination Types ---

type PaginatedList[T any] struct {
	Data       []T   `json:"data"`
	TotalCount int64 `json:"total_count"`
}
