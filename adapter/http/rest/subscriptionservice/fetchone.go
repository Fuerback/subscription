package subscriptionservice

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Fuerback/subscription/core/dto"
)

func (service service) FetchOne(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	id := strings.TrimPrefix(request.URL.Path, "/v1/subscription/")
	subscription, err := service.usecase.FetchOne(id)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	subscriptionDetails := dto.FromDomainToDtoSubscriptionDetails(*subscription)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(subscriptionDetails)
}
