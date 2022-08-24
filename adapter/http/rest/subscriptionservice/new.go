package subscriptionservice

import (
	"net/http"

	"github.com/Fuerback/subscription/core/usecase/subscriptionusecase"
)

type service struct {
	usecase subscriptionusecase.SubscriptionUseCase
}

// SubscriptionService is a contract of http adapter layer
type SubscriptionService interface {
	FetchOne(response http.ResponseWriter, request *http.Request)
	UpdateStatus(response http.ResponseWriter, request *http.Request)
}

// New returns contract implementation of ProductService
func New(usecase subscriptionusecase.SubscriptionUseCase) SubscriptionService {
	return &service{
		usecase: usecase,
	}
}
