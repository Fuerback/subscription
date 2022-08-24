package dto

import (
	"encoding/json"
	"io"

	"github.com/Fuerback/subscription/core/domain"
)

// CreateProductRequest is an representation request body to create a new Product
type CreateProductRequest struct {
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Period string  `json:"period"`
	Active string  `json:"active"`
}

// Product is entity of table product database column
type ProductResponse struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Price  float32 `json:"price,omitempty"`
	Period string  `json:"period,omitempty"`
	Active string  `json:"active,omitempty"`
}

// FromJSONCreateProductRequest converts json body request to a CreateProductRequest struct
func FromJSONCreateProductRequest(body io.Reader) (*CreateProductRequest, error) {
	createProductRequest := CreateProductRequest{}
	if err := json.NewDecoder(body).Decode(&createProductRequest); err != nil {
		return nil, err
	}

	return &createProductRequest, nil
}

func FromDomainToDtoProduct(product domain.Product) ProductResponse {
	return ProductResponse{
		ID:     product.ID,
		Name:   product.Name,
		Price:  product.Price,
		Period: product.Period,
		Active: product.Active,
	}
}
