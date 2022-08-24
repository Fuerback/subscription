package subscriptionrepository

import (
	"github.com/Fuerback/subscription/adapter/sqlite"
	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

type repository struct {
	db sqlite.Db
}

// SubscriptionRepository is a contract of database connection adapter layer
type SubscriptionRepository interface {
	FetchOne(id string) (*domain.SubscriptionDetails, error)
	UpdateStatus(id string, status *dto.UpdateSubscriptionStatus) error
}

// New returns contract implementation of ProductRepository
func New(db sqlite.Db) SubscriptionRepository {
	return &repository{
		db: db,
	}
}
