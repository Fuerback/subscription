package domain

import (
	"net/http"

	"github.com/Fuerback/subscription/core/dto"
)

// Product is entity of table product database column
type Product struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Price  float32 `json:"price,omitempty"`
	Period string  `json:"period,omitempty"`
	Active string  `json:"active,omitempty"`
}

const (
	Monthly string = "MONTHLY"
	Annual         = "ANNUAL"
)

// ProductService is a contract of http adapter layer
type ProductService interface {
	Fetch(response http.ResponseWriter, request *http.Request)
	FetchOne(response http.ResponseWriter, request *http.Request)
	Purchase(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type ProductUseCase interface {
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
	FetchOne(id string) (*Product, error)
	Purchase(purchaseRequest *dto.PurchaseRequest) (*Subscription, error)
}

// ProductRepository is a contract of database connection adapter layer
type ProductRepository interface {
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
	FetchOne(id string) (*Product, error)
	Purchase(subscription *Subscription) (string, error)
}
