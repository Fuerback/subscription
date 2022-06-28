package productrepository

import (
	"context"

	"github.com/Fuerback/subscription/core/domain"
)

func (repository repository) FetchOne(id string) (*domain.Product, error) {
	ctx := context.Background()
	product := &domain.Product{}

	rows := repository.db.QueryRowContext(ctx, "select * from product where id = ?", id)
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Period,
		&product.Price,
		&product.Active,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
