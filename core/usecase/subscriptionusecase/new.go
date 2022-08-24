package subscriptionusecase

import (
	"github.com/Fuerback/subscription/adapter/sqlite/subscriptionrepository"
	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

type usecase struct {
	repository subscriptionrepository.SubscriptionRepository
}

// SubscriptionUseCase is a contract of business rule layer
type SubscriptionUseCase interface {
	FetchOne(id string) (*domain.SubscriptionDetails, error)
	UpdateStatus(id string, status *dto.UpdateSubscriptionStatus) error
}

// New returns contract implementation of ProductUseCase
func New(repository subscriptionrepository.SubscriptionRepository) SubscriptionUseCase {
	return &usecase{
		repository: repository,
	}
}
