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
	Name           *string   `json:"name,omitempty"`
	Slug           *string   `json:"slug,omitempty"`
	Currency       *string   `json:"currency,omitempty"`
	SubscriberType *[]string `json:"subscriber_type,omitempty"`
}

type GetProductByIDParams struct {
	APIParams
	ID string `json:"id"`
}

type CommerceProduct struct {
	APIResource
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
	Currency        string          `json:"currency"`
	Plans           []*CommercePlan `json:"plans"`
	SubscriberType  []string        `json:"subscriber_type"`
	OwnerEntityType string          `json:"owner_entity_type"`
	CreatedAt       string          `json:"created_at"`
	UpdatedAt       string          `json:"updated_at"`
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

// --- Feature Types ---

type CommerceFeatureUnitPricing struct {
	Amount int64 `json:"amount"`
	Qty    int   `json:"qty"`
}

type CommerceFeature struct {
	APIResource
	ID              string                       `json:"id"`
	PlanID          string                       `json:"plan_id"`
	Name            string                       `json:"name"`
	Description     string                       `json:"description"`
	AvatarURL       string                       `json:"avatar_url"`
	Slug            string                       `json:"slug"`
	PubliclyVisible bool                         `json:"publicly_visible"`
	IncludeInJWT    bool                         `json:"include_in_jwt"`
	JWTValue        string                       `json:"jwt_value"`
	IsAddon         bool                         `json:"is_addon"`
	BaseFeeAmount   int64                        `json:"base_fee_amount"`
	IsMetered       bool                         `json:"is_metered"`
	BillingModel    string                       `json:"billing_model"`
	UnitName        string                       `json:"unit_name"`
	UnitNamePlural  string                       `json:"unit_name_plural"`
	HasTrialUnits   bool                         `json:"has_trial_units"`
	TrialUnits      int                          `json:"trial_units"`
	UnitPricing     []CommerceFeatureUnitPricing `json:"unit_pricing"`
	CreatedAt       string                       `json:"created_at"`
	UpdatedAt       string                       `json:"updated_at"`
}

type CommerceFeatureList PaginatedList[CommerceFeature]

type CreateFeatureParams struct {
	APIParams
	Name            string                       `json:"name"`
	Description     string                       `json:"description"`
	PlanID          string                       `json:"plan_id"`
	AvatarURL       string                       `json:"avatar_url"`
	Slug            string                       `json:"slug"`
	PubliclyVisible bool                         `json:"publicly_visible"`
	IncludeInJWT    bool                         `json:"include_in_jwt"`
	JWTValue        string                       `json:"jwt_value"`
	IsAddon         bool                         `json:"is_addon"`
	BaseFeeAmount   int64                        `json:"base_fee_amount"`
	IsMetered       bool                         `json:"is_metered"`
	BillingModel    string                       `json:"billing_model"`
	UnitName        string                       `json:"unit_name"`
	UnitNamePlural  string                       `json:"unit_name_plural"`
	TrialUnits      int                          `json:"trial_units"`
	UnitPricing     []CommerceFeatureUnitPricing `json:"unit_pricing"`
}

type UpdateFeatureParams struct {
	APIParams
	ID              string                        `json:"id"`
	Name            *string                       `json:"name,omitempty"`
	Description     *string                       `json:"description,omitempty"`
	AvatarURL       *string                       `json:"avatar_url,omitempty"`
	Slug            *string                       `json:"slug,omitempty"`
	PubliclyVisible *bool                         `json:"publicly_visible,omitempty"`
	IncludeInJWT    *bool                         `json:"include_in_jwt,omitempty"`
	JWTValue        *string                       `json:"jwt_value,omitempty"`
	IsAddon         *bool                         `json:"is_addon,omitempty"`
	BaseFeeAmount   *int64                        `json:"base_fee_amount,omitempty"`
	IsMetered       *bool                         `json:"is_metered,omitempty"`
	BillingModel    *string                       `json:"billing_model,omitempty"`
	UnitName        *string                       `json:"unit_name,omitempty"`
	UnitNamePlural  *string                       `json:"unit_name_plural,omitempty"`
	TrialUnits      *int                          `json:"trial_units,omitempty"`
	UnitPricing     *[]CommerceFeatureUnitPricing `json:"unit_pricing,omitempty"`
}

type ListFeaturesByInstanceIDParams struct {
	APIParams
	ID string `json:"id"`
}

type ListFeaturesByPlanIDParams struct {
	APIParams
	ID string `json:"id"`
}

// --- Plan Types ---

type CreatePlanParams struct {
	APIParams
	Name        string `json:"name"`
	ProductID   string `json:"product_id"`
	Amount      int64  `json:"amount"`
	IsRecurring bool   `json:"is_recurring"`
	IsProrated  bool   `json:"is_prorated"`
	Period      string `json:"period"`
	Interval    int    `json:"interval"`
	AvatarURL   string `json:"avatar_url"`
	Description string `json:"description"`
}

type UpdatePlanParams struct {
	APIParams
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Amount      *int64  `json:"amount,omitempty"`
	IsRecurring *bool   `json:"is_recurring,omitempty"`
	IsProrated  *bool   `json:"is_prorated,omitempty"`
	Period      *string `json:"period,omitempty"`
	Interval    *int    `json:"interval,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
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
	Amount          int64            `json:"amount"`
	IsRecurring     bool             `json:"is_recurring"`
	IsProrated      bool             `json:"is_prorated"`
	Period          string           `json:"period"`
	Interval        int              `json:"interval"`
	AvatarURL       string           `json:"avatar_url"`
	ProductID       string           `json:"product_id"`
	Description     string           `json:"description"`
	Slug            string           `json:"slug"`
	BillingCycles   *int             `json:"billing_cycles,omitempty"`
	SubscriberCount int64            `json:"subscriber_count"`
	CreatedAt       string           `json:"created_at"`
	UpdatedAt       string           `json:"updated_at"`
}

type CommercePlanWithNoProduct struct {
	CommercePlan
	Product *CommerceProduct `json:"-"`
}

type CommercePlanList PaginatedList[CommercePlan]

type ListPlansByInstanceIDParams struct {
	APIParams
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
