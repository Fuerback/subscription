package subscriptionusecase

import "github.com/Fuerback/subscription/core/domain"

type usecase struct {
	repository domain.SubscriptionRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.SubscriptionRepository) domain.SubscriptionUseCase {
	return &usecase{
		repository: repository,
	}
}
