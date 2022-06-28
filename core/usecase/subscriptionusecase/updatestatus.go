package subscriptionusecase

import (
	"github.com/Fuerback/subscription/core/dto"
)

func (usecase usecase) UpdateStatus(id string, status *dto.UpdateSubscriptionStatus) error {
	err := usecase.repository.UpdateStatus(id, status)

	if err != nil {
		return err
	}

	return nil
}
