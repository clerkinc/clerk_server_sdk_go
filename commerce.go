package clerk

import "time"

// CommercePlan represents a subscription plan.
type CommercePlan struct {
	APIResource
	AppID           string          `json:"app_id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
	ImageURL        string          `json:"image_url"`
	Description     *string         `json:"description,omitempty"`
	Product         CommerceProduct `json:"product"`
	BaseAmount      int64           `json:"base_amount"`
	IsRecurring     bool            `json:"is_recurring"`
	IsProrated      bool            `json:"is_prorated"`
	Period          string          `json:"period"` // Enum equivalent for CommercePlanPeriod
	Interval        int             `json:"interval"`
	BillingCycles   *int            `json:"billing_cycles,omitempty"`
	SubscriberCount int64           `json:"subscriber_count"`
}

// CommerceProduct represents a product associated with a plan.
type CommerceProduct struct {
	APIResource
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Currency        string    `json:"currency"`
	SubscriberType  []string  `json:"subscriber_type"`   // Enum equivalent for CommerceProductSubscriberType
	OwnerEntityType string    `json:"owner_entity_type"` // Enum equivalent for CommerceProductOwnerEntity
}

// CommerceProductWithPlans combines a product and its associated plans.
type CommerceProductWithPlans struct {
	CommerceProduct
	Plans []CommercePlan `json:"plans"`
}

// CommerceCustomer represents a customer subscribing to a product.
type CommerceCustomer struct {
	APIResource
	AppID      string    `json:"app_id"`
	CreatedAt  time.Time `json:"created_at"`
	EntityType string    `json:"entity_type"` // Enum equivalent for CommerceProductSubscriberType
	Entity     struct {
		ID       string  `json:"id"`
		Name     string  `json:"name"`
		ImageURL *string `json:"image_url,omitempty"`
	} `json:"entity"`
}

// CommerceSubscription represents a subscription.
type CommerceSubscription struct {
	APIResource
	AppID           string           `json:"app_id"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Customer        CommerceCustomer `json:"customer"`
	Plan            CommercePlan     `json:"plan"`
	Status          string           `json:"status"` // Enum equivalent for CommerceSubscriptionStatus
	ToBeCancelledAt *time.Time       `json:"to_be_cancelled_at,omitempty"`
	LastInvoice     *InvoiceSummary  `json:"last_invoice,omitempty"`
	NextInvoice     *InvoiceSummary  `json:"next_invoice,omitempty"`
}

// InvoiceSummary is a reduced representation of an invoice.
type InvoiceSummary struct {
	APIResource
	DueAt  time.Time `json:"due_at"`
	Amount int64     `json:"amount"`
	Status string    `json:"status"` // Enum equivalent for CommerceInvoiceStatus
}

// CommerceInvoice represents a detailed invoice.
type CommerceInvoice struct {
	APIResource
	AppID                    string               `json:"app_id"`
	CreatedAt                time.Time            `json:"created_at"`
	UpdatedAt                time.Time            `json:"updated_at"`
	Subscription             CommerceSubscription `json:"subscription"`
	Amount                   int64                `json:"amount"`
	Status                   string               `json:"status"` // Enum equivalent for CommerceInvoiceStatus
	DueAt                    time.Time            `json:"due_at"`
	FinalizingPaymentAttempt string               `json:"finalizing_payment_attempt_id"`
}

// CommercePaymentAttempt represents a payment attempt for an invoice.
type CommercePaymentAttempt struct {
	APIResource
	AppID     string           `json:"app_id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Customer  CommerceCustomer `json:"customer"`
	Invoice   CommerceInvoice  `json:"invoice"`
	Status    string           `json:"status"` // Enum equivalent for CommercePaymentAttemptStatus
	Amount    int64            `json:"amount"`
}

// PaginatedList is a generic response for paginated resources.
type PaginatedList[T any] struct {
	APIResource
	Data       []T   `json:"data"`
	TotalCount int64 `json:"total_count"`
}

// ListCommerceSubscriptionsResponse represents a paginated list of subscriptions.
type ListCommerceSubscriptionsResponse struct {
	PaginatedList[CommerceSubscription]
}

// ListCommerceInvoicesResponse represents a paginated list of invoices.
type ListCommerceInvoicesResponse struct {
	PaginatedList[CommerceInvoice]
}

// ListCommercePaymentAttemptsResponse represents a paginated list of payment attempts.
type ListCommercePaymentAttemptsResponse struct {
	PaginatedList[CommercePaymentAttempt]
}
