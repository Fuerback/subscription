package productusecase

import (
	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) ([]domain.Product, error) {
	products, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
