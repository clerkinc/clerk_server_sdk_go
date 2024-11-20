package clerk

import "time"

// --- Product Types ---

type CreateProductParams struct {
	APIParams
	InstanceID      *string   `json:"instance_id,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Slug            *string   `json:"slug,omitempty"`
	Currency        *string   `json:"currency,omitempty"`
	SubscriberType  *[]string `json:"subscriber_type,omitempty"`
	OwnerEntityType *string   `json:"owner_entity_type,omitempty"`
}

type UpdateProductParams struct {
	APIParams
	ID              *string   `json:"id,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Slug            *string   `json:"slug,omitempty"`
	Currency        *string   `json:"currency,omitempty"`
	SubscriberType  *[]string `json:"subscriber_type,omitempty"`
	OwnerEntityType *string   `json:"owner_entity_type,omitempty"`
}

type GetProductByIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type CommerceProduct struct {
	APIResource
	Name            *string   `json:"name,omitempty"`
	Slug            *string   `json:"slug,omitempty"`
	Currency        *string   `json:"currency,omitempty"`
	SubscriberType  *[]string `json:"subscriber_type,omitempty"`
	OwnerEntityType *string   `json:"owner_entity_type,omitempty"`
}

type CommerceProductWithPlans struct {
	CommerceProduct
	Plans *[]CommercePlan `json:"plans,omitempty"`
}

type ListProductsByInstanceIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type ListCommerceProductsResponse struct {
	APIParams
	PaginatedList[CommerceProduct]
}

// --- Plan Types ---

type CreatePlanParams struct {
	APIParams
	Name        *string `json:"name,omitempty"`
	ProductID   *string `json:"product_id,omitempty"`
	BaseAmount  *int64  `json:"base_amount,omitempty"`
	IsRecurring *bool   `json:"is_recurring,omitempty"`
}

type UpdatePlanParams struct {
	APIParams
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetPlanByIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type CommercePlan struct {
	APIResource
	Name            *string `json:"name,omitempty"`
	ProductID       *string `json:"product_id,omitempty"`
	BaseAmount      *int64  `json:"base_amount,omitempty"`
	IsRecurring     *bool   `json:"is_recurring,omitempty"`
	Period          *string `json:"period,omitempty"`
	Interval        *int    `json:"interval,omitempty"`
	BillingCycles   *int    `json:"billing_cycles,omitempty"`
	SubscriberCount *int64  `json:"subscriber_count,omitempty"`
}

type ListPlansByInstanceIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

// --- Integration Types ---

type CreateIntegrationParams struct {
	APIParams
	InstanceID *string `json:"instance_id,omitempty"`
	Email      *string `json:"email,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type UpdateIntegrationParams struct {
	APIParams
	CommerceIntegrationID *string `json:"id,omitempty"`
	Status                *string `json:"status,omitempty"`
}

type GetIntegrationParams struct {
	APIParams
	IntegrationID *string `json:"id,omitempty"`
}

type CommerceIntegration struct {
	APIResource
	IntegrationID   *string `json:"integration_id,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty"`
	Status          *string `json:"status,omitempty"`
}

type CommerceIntegrationResponse struct {
	APIResource
	URL *string `json:"url,omitempty"`
}

type ListCommerceIntegrationsResponse struct {
	APIParams
	PaginatedList[CommerceIntegration]
}

type ListIntegrationsByInstanceIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

// --- Subscription Types ---

type CreateSubscriptionParams struct {
	APIParams
	CustomerID *string `json:"customer_id,omitempty"`
	PlanID     *string `json:"plan_id,omitempty"`
	Status     *string `json:"status,omitempty"`
}

type UpdateSubscriptionParams struct {
	APIParams
	ID     *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

type GetSubscriptionByIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type ListSubscriptionsByInstanceIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type ListSubscriptionsByUserIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type CommerceSubscription struct {
	APIResource
	Customer *CommerceCustomer `json:"customer,omitempty"`
	Plan     *CommercePlan     `json:"plan,omitempty"`
	Status   *string           `json:"status,omitempty"`
}

type ListCommerceSubscriptionsResponse struct {
	APIParams
	PaginatedList[CommerceSubscription]
}

// --- Invoice Types ---

type CreateInvoiceParams struct {
	APIParams
	SubscriptionID *string `json:"subscription_id,omitempty"`
	Amount         *int64  `json:"amount,omitempty"`
	DueAt          *string `json:"due_at,omitempty"`
}

type UpdateInvoiceParams struct {
	APIParams
	ID     *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

type GetInvoiceByIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type ListInvoicesByInstanceIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type CommerceInvoice struct {
	APIResource
	Subscription *CommerceSubscription `json:"subscription,omitempty"`
	Amount       *int64                `json:"amount,omitempty"`
	Status       *string               `json:"status,omitempty"`
	DueAt        *time.Time            `json:"due_at,omitempty"`
}

type ListCommerceInvoicesResponse struct {
	APIParams
	PaginatedList[CommerceInvoice]
}

// --- Payment Attempt Types ---

type CreatePaymentAttemptParams struct {
	APIParams
	InvoiceID *string `json:"invoice_id,omitempty"`
	Amount    *int64  `json:"amount,omitempty"`
	Status    *string `json:"status,omitempty"`
}

type UpdatePaymentAttemptParams struct {
	APIParams
	ID     *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

type GetPaymentAttemptByIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type ListPaymentAttemptsByInstanceIDParams struct {
	APIParams
	ID *string `json:"id,omitempty"`
}

type CommercePaymentAttempt struct {
	APIResource
	Invoice *CommerceInvoice `json:"invoice,omitempty"`
	Amount  *int64           `json:"amount,omitempty"`
	Status  *string          `json:"status,omitempty"`
}

type ListCommercePaymentAttemptsResponse struct {
	APIParams
	PaginatedList[CommercePaymentAttempt]
}

// --- Customer Types ---

type CommerceCustomer struct {
	Entity *struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"entity,omitempty"`
}

// --- Pagination Types ---

type PaginatedList[T any] struct {
	APIResource
	Data       *[]T   `json:"data,omitempty"`
	TotalCount *int64 `json:"total_count,omitempty"`
}
