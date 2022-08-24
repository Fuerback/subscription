package domain

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
	PausedAt     string  `json:"paused_at,omitempty"`
	CancelledAt  string  `json:"cancelled_at,omitempty"`
	Product      Product `json:"product,omitempty"`
	Account      Account `json:"account,omitempty"`
	Status       string  `json:"status,omitempty"`
	Voucher      string  `json:"voucher,omitempty"`
	PaymentValue float32 `json:"payment_value,omitempty"`
}
