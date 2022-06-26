package productusecase

import (
	"github.com/Fuerback/subscription/core/domain"
)

func (usecase usecase) FetchOne(id string) (*domain.Product, error) {
	product, err := usecase.repository.FetchOne(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
