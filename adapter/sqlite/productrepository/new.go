package productrepository

import (
	"github.com/Fuerback/subscription/adapter/sqlite"
	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

type repository struct {
	db sqlite.Db
}

// ProductRepository is a contract of database connection adapter layer
type ProductRepository interface {
	Fetch(paginationRequest *dto.PaginationRequestParms) ([]domain.Product, error)
	FetchOne(id string) (*domain.Product, error)
	Purchase(subscription *domain.Subscription) (string, error)
}

// New returns contract implementation of ProductRepository
func New(db sqlite.Db) ProductRepository {
	return &repository{
		db: db,
	}
}
