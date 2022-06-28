package domain

import "net/http"

const (
	Active    string = "ACTIVE"
	Paused           = "PAUSED"
	Cancelled        = "CANCELLED"
	Trial            = "TRIAL"
)

// Subscription is entity of table subscription database column
type Subscription struct {
	ID           string  `json:"subscription_id"`
	StartsAt     string  `json:"starts_at,omitempty"`
	EndsAt       string  `json:"ends_at,omitempty"`
	Product      string  `json:"product,omitempty"`
	Account      string  `json:"account,omitempty"`
	Status       string  `json:"status,omitempty"`
	Voucher      string  `json:"voucher,omitempty"`
	PaymentValue float32 `json:"payment_value,omitempty"`
}

// SubscriptionDetails is entity of table subscription database column
type SubscriptionDetails struct {
	ID           string  `json:"subscription_id"`
	StartsAt     string  `json:"starts_at,omitempty"`
	EndsAt       string  `json:"ends_at,omitempty"`
	Product      Product `json:"product,omitempty"`
	Account      Account `json:"account,omitempty"`
	Status       string  `json:"status,omitempty"`
	Voucher      string  `json:"voucher,omitempty"`
	PaymentValue float32 `json:"payment_value,omitempty"`
}

// SubscriptionService is a contract of http adapter layer
type SubscriptionService interface {
	FetchOne(response http.ResponseWriter, request *http.Request)
}

// SubscriptionUseCase is a contract of business rule layer
type SubscriptionUseCase interface {
	FetchOne(id string) (*SubscriptionDetails, error)
}

// SubscriptionRepository is a contract of database connection adapter layer
type SubscriptionRepository interface {
	FetchOne(id string) (*SubscriptionDetails, error)
}
