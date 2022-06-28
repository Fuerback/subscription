package subscriptionusecase

import "github.com/Fuerback/subscription/core/domain"

func (usecase usecase) FetchOne(id string) (*domain.SubscriptionDetails, error) {
	subscription, err := usecase.repository.FetchOne(id)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}
