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
	APIResource
	CommerceProduct
	Plans []CommercePlan `json:"plans"`
}

type ListProductsByInstanceIDParams struct {
	APIParams
	ID string `json:"id"`
}

type CommerceProductList PaginatedList[CommerceProduct]

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
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Product         *CommerceProduct `json:"product,omitempty"`
	BaseAmount      int64            `json:"base_amount"`
	IsRecurring     bool             `json:"is_recurring"`
	IsProrated      bool             `json:"is_prorated"`
	Period          string           `json:"period"`
	Interval        int              `json:"interval"`
	BillingCycles   *int             `json:"billing_cycles,omitempty"`
	SubscriberCount int64            `json:"subscriber_count"`
	CreatedAt       string           `json:"created_at"`
	UpdatedAt       string           `json:"updated_at"`
}

type CommercePlanWithNoProduct struct {
	CommercePlan
	Product *CommerceProduct `json:"-"`
}

type ListPlansByInstanceIDParams struct {
	APIParams
}

// --- Integration Types ---
type CommerceIntegration struct {
	APIResource
	ID                string `json:"id"`
	IntegrationType   string `json:"integration_type"`
	IntegrationStatus string `json:"integration_status"`
	CreatedAt         string `json:"created_at"` // ISO 8601 format
	UpdatedAt         string `json:"updated_at"` // ISO 8601 format
	InstanceID        string `json:"instance_id"`
	URL               string `json:"url"`
}

type CommerceIntegrationList struct {
	APIResource
	PaginatedList[CommerceIntegration]
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
	CustomerID      string `json:"customer_id"`
	PlanID          string `json:"plan_id"`
	PaymentSourceID string `json:"payment_source_id"`
}

type UpdateSubscriptionParams struct {
	APIParams
	Status *string `json:"status,omitempty"`
}

type GetSubscriptionByIDParams struct {
	APIParams
	ID string `json:"id"`
}

type ListSubscriptionsByUserIDParams struct {
	APIParams
	ID             string `json:"id"`
	SubscriberType string `json:"subscriber_type"`
}

type CommerceSubscription struct {
	APIResource
	ID              string            `json:"id"`
	AppID           string            `json:"app_id"`
	Customer        *CommerceCustomer `json:"customer,omitempty"`
	InstanceID      string            `json:"instance_id"`
	PaymentSourceID string            `json:"payment_source_id"`
	PlanID          string            `json:"plan_id"`
	Plan            *CommercePlan     `json:"plan,omitempty"`
	Status          string            `json:"status"`
	LastInvoice     *CommerceInvoice  `json:"last_invoice,omitempty"`
	NextInvoice     *CommerceInvoice  `json:"next_invoice,omitempty"`
	CreatedAt       string            `json:"created_at"` // ISO 8601 format
	UpdatedAt       string            `json:"updated_at"` // ISO 8601 format
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
type CommerceInvoice struct {
	APIResource
	ID                       string                `json:"id"`
	Subscription             *CommerceSubscription `json:"subscription,omitempty"`
	Amount                   int64                 `json:"amount"`
	Status                   string                `json:"status"`
	DueAt                    *time.Time            `json:"due_at,omitempty"`
	FinalizingPaymentAttempt string                `json:"finalizing_payment_attempt_id,omitempty"`
}

type CommerceInvoiceList PaginatedList[CommerceInvoice]

// --- Payment Attempt Types ---

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
