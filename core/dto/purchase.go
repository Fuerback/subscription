package dto

import (
	"encoding/json"
	"net/http"
	"strings"
)

// PurchaseRequest is an representation request body to purchase a new Product
type PurchaseRequest struct {
	ProductID string
	AccountID string
	Voucher   string `json:"voucher"`
}

// FromJSONPurchaseProductRequest converts json body request to a PurchaseRequest struct
func FromJSONPurchaseProductRequest(request *http.Request) (*PurchaseRequest, error) {
	purchaseProductRequest := PurchaseRequest{}
	if err := json.NewDecoder(request.Body).Decode(&purchaseProductRequest); err != nil {
		return nil, err
	}

	purchaseProductRequest.AccountID = request.Header.Get("x-account-id")
	purchaseProductRequest.ProductID = strings.TrimPrefix(request.URL.Path, "/v1/product/purchase/")

	return &purchaseProductRequest, nil
}
