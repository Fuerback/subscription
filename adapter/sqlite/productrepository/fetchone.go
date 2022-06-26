package productrepository

import (
	"context"

	"github.com/Fuerback/subscription/core/domain"
)

func (repository repository) FetchOne(id string) (*domain.Product, error) {
	ctx := context.Background()
	product := &domain.Product{}

	stmt, err := repository.db.PrepareContext(ctx, "select * from product where id = ?")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Period, &product.Price, &product.Active)
	if err != nil {
		return nil, err
	}

	return product, nil
}
