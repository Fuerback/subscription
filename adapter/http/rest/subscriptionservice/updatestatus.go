package subscriptionservice

import (
	"net/http"
	"strings"

	"github.com/Fuerback/subscription/core/dto"
)

func (service service) UpdateStatus(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	updateStatus, err := dto.FromJSONUpdateSubscriptionStatusRequest(request.Body)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	id := strings.TrimPrefix(request.URL.Path, "/v1/subscription/status/")
	err = service.usecase.UpdateStatus(id, updateStatus)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	response.WriteHeader(http.StatusOK)
}
