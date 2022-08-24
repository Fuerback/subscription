package dto

import (
	"encoding/json"
	"io"

	"github.com/Fuerback/subscription/core/domain"
)

// UpdateSubscriptionStatus is an representation request body to update subscription status
type UpdateSubscriptionStatus struct {
	Status string `json:"status" validate:"required,oneof=PAUSED ACTIVE"`
}

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

// SubscriptionDetails is an representation of entity subscription
type SubscriptionDetailsResponse struct {
	ID           string          `json:"subscription_id"`
	StartsAt     string          `json:"starts_at,omitempty"`
	EndsAt       string          `json:"ends_at,omitempty"`
	PausedAt     string          `json:"paused_at,omitempty"`
	CancelledAt  string          `json:"cancelled_at,omitempty"`
	Product      ProductResponse `json:"product,omitempty"`
	Account      AccountResponse `json:"account,omitempty"`
	Status       string          `json:"status,omitempty"`
	Voucher      string          `json:"voucher,omitempty"`
	PaymentValue float32         `json:"payment_value,omitempty"`
}

// FromJSONUpdateSubscriptionStatusRequest converts json body request to a UpdateSubscriptionStatus struct
func FromJSONUpdateSubscriptionStatusRequest(body io.ReadCloser) (*UpdateSubscriptionStatus, error) {
	updatestatusRequest := UpdateSubscriptionStatus{}
	if err := json.NewDecoder(body).Decode(&updatestatusRequest); err != nil {
		return nil, err
	}

	return &updatestatusRequest, nil
}

func FromDomainToDtoSubscription(subscription domain.Subscription) Subscription {
	return Subscription{
		ID:           subscription.ID,
		StartsAt:     subscription.StartsAt,
		EndsAt:       subscription.EndsAt,
		Product:      subscription.Product,
		Account:      subscription.Account,
		Status:       subscription.Status,
		Voucher:      subscription.Voucher,
		PaymentValue: subscription.PaymentValue,
	}
}

func FromDomainToDtoSubscriptionDetails(subscription domain.SubscriptionDetails) SubscriptionDetailsResponse {
	return SubscriptionDetailsResponse{
		ID:           subscription.ID,
		StartsAt:     subscription.StartsAt,
		EndsAt:       subscription.EndsAt,
		PausedAt:     subscription.PausedAt,
		CancelledAt:  subscription.CancelledAt,
		Product:      FromDomainToDtoProduct(subscription.Product),
		Account:      FromDomainToDtoAccount(subscription.Account),
		Status:       subscription.Status,
		Voucher:      subscription.Voucher,
		PaymentValue: subscription.PaymentValue,
	}
}
