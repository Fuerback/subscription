package domain

const (
	Active    string = "ACTIVE"
	Paused           = "PAUSED"
	Cancelled        = "CANCELLED"
	Trial            = "TRIAL"
)

// Subscription is entity of table subscription database column
type Subscription struct {
	ID           string
	StartsAt     string
	EndsAt       string
	Product      string
	Account      string
	Status       string
	Voucher      string
	PaymentValue float32
}

// SubscriptionDetails is entity of table subscription database column
type SubscriptionDetails struct {
	ID           string
	StartsAt     string
	EndsAt       string
	PausedAt     string
	CancelledAt  string
	Product      Product
	Account      Account
	Status       string
	Voucher      string
	PaymentValue float32
}
