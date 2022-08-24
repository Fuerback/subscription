package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/Fuerback/subscription/core/dto"
)

func (service service) Purchase(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	purchaseRequest, err := dto.FromJSONPurchaseProductRequest(request)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	subscription, err := service.usecase.Purchase(purchaseRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	subscriptionResponse := dto.FromDomainToDtoSubscription(*subscription)

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(subscriptionResponse)
}
