package domain

// Product is entity of table product database column
type Product struct {
	ID     string
	Name   string
	Price  float32
	Period string
	Active string
}

const (
	Monthly string = "MONTHLY"
	Annual         = "ANNUAL"
)
