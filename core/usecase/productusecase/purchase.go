package productusecase

import (
	"time"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

func (usecase usecase) Purchase(purchaseRequest *dto.PurchaseRequest) (*domain.Subscription, error) {
	product, err := usecase.FetchOne(purchaseRequest.ProductID)
	if err != nil {
		return nil, err
	}

	subscription := &domain.Subscription{
		StartsAt:     time.Now().Format("2006-02-01"),
		EndsAt:       getEndsAtPeriod(product.Period),
		Product:      purchaseRequest.ProductID,
		Account:      purchaseRequest.AccountD,
		Status:       domain.Active,
		PaymentValue: product.Price,
		Voucher:      purchaseRequest.Voucher,
	}

	id, err := usecase.repository.Purchase(subscription)
	if err != nil {
		return nil, err
	}

	subscription.ID = id

	return subscription, nil
}

func getEndsAtPeriod(period string) string {
	now := time.Now()
	if period == domain.Monthly {
		return now.Add(time.Hour * 24 * 30).Format("2006-02-01")
	}
	return now.Add(time.Hour * 24 * 365).Format("2006-02-01")
}
