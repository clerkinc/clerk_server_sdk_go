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
	PayerType       []string `json:"payer_type"`
	OwnerEntityType string   `json:"owner_entity_type"`
}

type UpdateProductParams struct {
	APIParams
	Name      *string   `json:"name,omitempty"`
	Slug      *string   `json:"slug,omitempty"`
	Currency  *string   `json:"currency,omitempty"`
	PayerType *[]string `json:"payer_type,omitempty"`
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
	PayerType       []string        `json:"payer_type"`
	OwnerEntityType string          `json:"owner_entity_type"`
	CreatedAt       int64           `json:"created_at"`
	UpdatedAt       int64           `json:"updated_at"`
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

type CommercePlanFeature struct {
	APIResource
	ID        string `json:"id"`
	PlanID    string `json:"plan_id"`
	FeatureID string `json:"feature_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CommercePlanFeatureList PaginatedList[CommercePlanFeature]

type CommerceFeature struct {
	APIResource
	ID              string                       `json:"id"`
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
	CreatedAt       int64                        `json:"created_at"`
	UpdatedAt       int64                        `json:"updated_at"`
}

type CommerceFeatureList PaginatedList[CommerceFeature]

type CreatePlanFeatureParams struct {
	APIParams
	PlanID    string `json:"plan_id"`
	FeatureID string `json:"feature_id"`
}

type CreateMultiplePlanFeaturesParams struct {
	APIParams
	PlanID     string   `json:"plan_id"`
	FeatureIDs []string `json:"feature_ids"`
}

type DeletePlanFeaturesParams struct {
	APIParams
	FeatureIDs []string `json:"feature_ids"`
	PlanID     string   `json:"plan_id"`
}

type CreateFeatureParams struct {
	APIParams
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
	TrialUnits      int                          `json:"trial_units"`
	UnitPricing     []CommerceFeatureUnitPricing `json:"unit_pricing"`
}

type CreateMultipleFeaturesParams struct {
	APIParams
	Features []CreateFeatureParams `json:"features"`
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
	Slug        string `json:"slug"`
	Amount      int64  `json:"amount"`
	IsRecurring bool   `json:"is_recurring"`
	IsProrated  bool   `json:"is_prorated"`
	IsFree			bool	 `json:"is_free"`
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
	Slug        *string `json:"slug,omitempty"`
	IsRecurring *bool   `json:"is_recurring,omitempty"`
	Description *string `json:"description,omitempty"`
	IsProrated  *bool   `json:"is_prorated,omitempty"`
	IsFree			bool	 	`json:"is_free"`
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
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Product       *CommerceProduct  `json:"product,omitempty"`
	Amount        int64             `json:"amount"`
	IsRecurring   bool              `json:"is_recurring"`
	IsProrated    bool              `json:"is_prorated"`
	IsFree				bool	 						`json:"is_free"`
	IsDefault			bool							`json:"is_default"`
	Period        string            `json:"period"`
	Interval      int               `json:"interval"`
	AvatarURL     string            `json:"avatar_url"`
	ProductID     string            `json:"product_id"`
	Description   string            `json:"description"`
	Slug          string            `json:"slug"`
	BillingCycles *int              `json:"billing_cycles,omitempty"`
	PayerCount    int64             `json:"payer_count"`
	CreatedAt     int64             `json:"created_at"`
	UpdatedAt     int64             `json:"updated_at"`
	Features      []CommerceFeature `json:"features"`
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
	PayerID         string `json:"payer_id"`
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
	ID        string `json:"id"`
	PayerType string `json:"payer_type"`
}

type CommerceSubscription struct {
	APIResource
	ID              string           `json:"id"`
	AppID           string           `json:"app_id"`
	Payer           *CommercePayer   `json:"payer,omitempty"`
	InstanceID      string           `json:"instance_id"`
	PaymentSourceID string           `json:"payment_source_id"`
	PlanID          string           `json:"plan_id"`
	Plan            *CommercePlan    `json:"plan,omitempty"`
	Status          string           `json:"status"`
	LastInvoice     *CommerceInvoice `json:"last_invoice,omitempty"`
	NextInvoice     *CommerceInvoice `json:"next_invoice,omitempty"`
	CreatedAt       int64            `json:"created_at"` // ISO 8601 format
	UpdatedAt       int64            `json:"updated_at"` // ISO 8601 format
}

type ListCommerceSubscriptionsResponse struct {
	APIResource
	PaginatedList[CommerceSubscription]
}

type CreatePaymentSourceParams struct {
	APIParams
	PayerID    string `json:"payer_id"`
	Gateway    string `json:"gateway"`
	PayeeID    string `json:"payee_id"`
	ExternalID string `json:"external_id"`
	Last4      string `json:"last4"`
	CardType   string `json:"card_type"`
}

type CommercePaymentSource struct {
	APIResource
	ID         string `json:"id"`
	PayerID    string `json:"payer_id"`
	Gateway    string `json:"gateway"`
	PayeeID    string `json:"payee_id"`
	ExternalID string `json:"external_id"`
	CardType   string `json:"card_type"`
	LastFour   string `json:"last4"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type CommercePaymentSourceList PaginatedList[CommercePaymentSource]

// --- Subscription Types ---

type GetSubscriptionParams struct {
	APIParams
	ID string `json:"id"`
}

type ListPayersParams struct {
	APIParams
	InstanceID string `json:"instance_id"`
}

type CommercePayer struct {
	APIResource
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CommercePayerList PaginatedList[CommercePayer]

type CreatePayerParams struct {
	APIParams
	InstanceID string `json:"instance_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type ListCommercePayersResponse struct {
	APIResource
	PaginatedList[CommercePayer]
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

// --- Payee Types ---

type CommercePayee struct {
	APIResource

	ID            string `json:"id"`
	GatewayStatus string `json:"gateway_status"`
	GatewayType   string `json:"gateway_type"`
	StripeURL     string `json:"stripe_url"`
	StripeID      string `json:"stripe_id"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type CommercePayeeList PaginatedList[CommercePayee]