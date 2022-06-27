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
	StartsAt     string  `json:"starts_at"`
	EndsAt       string  `json:"ends_at"`
	Product      string  `json:"product"`
	Account      string  `json:"account"`
	Status       string  `json:"status"`
	Voucher      string  `json:"voucher"`
	PaymentValue float32 `json:"payment_value"`
}
