package subscriptionusecase

import (
	"fmt"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

func (usecase usecase) UpdateStatus(id string, status *dto.UpdateSubscriptionStatus) error {
	subscription, err := usecase.repository.FetchOne(id)
	if err != nil {
		return err
	}

	if subscription.Status == domain.Cancelled {
		return fmt.Errorf("Subscription cancelled, it is not posible to change the status")
	}

	if subscription.Status == status.Status {
		return nil
	}

	err = usecase.repository.UpdateStatus(id, status)

	if err != nil {
		return err
	}

	return nil
}
