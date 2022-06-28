package productservice

import (
	"encoding/json"
	"net/http"
	"strings"
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

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(product)
}
