package productservice

import (
	"net/http"

	"github.com/Fuerback/subscription/core/usecase/productusecase"
	"github.com/go-playground/validator/v10"
)

type service struct {
	usecase      productusecase.ProductUseCase
	jsonValidate *validator.Validate
}

// ProductService is a contract of http adapter layer
type ProductService interface {
	Fetch(response http.ResponseWriter, request *http.Request)
	FetchOne(response http.ResponseWriter, request *http.Request)
	Purchase(response http.ResponseWriter, request *http.Request)
}

// New returns contract implementation of ProductService
func New(usecase productusecase.ProductUseCase, jsonValidate *validator.Validate) ProductService {
	return &service{
		usecase:      usecase,
		jsonValidate: jsonValidate,
	}
}
