package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// PurchaseRequest is an representation request body to purchase a new Product
type PurchaseRequest struct {
	ProductID string `json:"-"`
	AccountID string `json:"-"`
	Voucher   string `json:"voucher,omitempty" validate:"omitempty,min=3,max=50"`
}

// FromJSONPurchaseProductRequest converts json body request to a PurchaseRequest struct
func FromJSONPurchaseProductRequest(request *http.Request) (*PurchaseRequest, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	purchaseProductRequest := PurchaseRequest{}
	if len(body) > 0 {
		err = json.Unmarshal(body, &purchaseProductRequest)
		if err != nil {
			return nil, fmt.Errorf("Could not unmarshal request: %s", err)
		}
	}

	purchaseProductRequest.AccountID = request.Header.Get("x-account-id")
	purchaseProductRequest.ProductID = strings.TrimPrefix(request.URL.Path, "/v1/product/purchase/")

	return &purchaseProductRequest, nil
}
