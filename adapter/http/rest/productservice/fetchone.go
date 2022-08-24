package productservice

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Fuerback/subscription/core/dto"
)

func (service service) FetchOne(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	id := strings.TrimPrefix(request.URL.Path, "/v1/product/")
	product, err := service.usecase.FetchOne(id)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	productResponse := dto.FromDomainToDtoProduct(*product)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(productResponse)
}
