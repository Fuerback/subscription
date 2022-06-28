package subscriptionservice

import "github.com/Fuerback/subscription/core/domain"

type service struct {
	usecase domain.SubscriptionUseCase
}

// New returns contract implementation of ProductService
func New(usecase domain.SubscriptionUseCase) domain.SubscriptionService {
	return &service{
		usecase: usecase,
	}
}
