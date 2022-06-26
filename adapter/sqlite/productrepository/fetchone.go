package productrepository

import (
	"context"

	"github.com/Fuerback/subscription/core/domain"
)

func (repository repository) FetchOne(id string) (*domain.Product, error) {
	ctx := context.Background()
	product := &domain.Product{}

	rows, err := repository.db.QueryContext(ctx, "select * from product where id = ?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Period,
			&product.Price,
			&product.Active,
		)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}
