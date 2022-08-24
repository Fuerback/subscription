package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/Fuerback/subscription/core/dto"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	paginationRequest, _ := dto.FromValuePaginationRequestParams(request)

	products, err := service.usecase.Fetch(paginationRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	paginationResponse := &dto.PaginationResponse{
		Items: products,
		Total: int32(len(products)),
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(paginationResponse)
}
