package productusecase

import (
	"github.com/Fuerback/subscription/adapter/sqlite/productrepository"
	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

type usecase struct {
	repository productrepository.ProductRepository
}

// ProductUseCase is a contract of business rule layer
type ProductUseCase interface {
	Fetch(paginationRequest *dto.PaginationRequestParms) ([]domain.Product, error)
	FetchOne(id string) (*domain.Product, error)
	Purchase(purchaseRequest *dto.PurchaseRequest) (*domain.Subscription, error)
}

// New returns contract implementation of ProductUseCase
func New(repository productrepository.ProductRepository) ProductUseCase {
	return &usecase{
		repository: repository,
	}
}
