package productrepository

import (
	"github.com/Fuerback/subscription/adapter/sqlite"
	"github.com/Fuerback/subscription/core/domain"
)

type repository struct {
	db sqlite.Db
}

// New returns contract implementation of ProductRepository
func New(db sqlite.Db) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
