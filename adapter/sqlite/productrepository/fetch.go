package productrepository

import (
	"context"
	"fmt"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination, error) {
	ctx := context.Background()
	products := []domain.Product{}

	query := fmt.Sprintf("select * from product limit %d offset %d", pagination.PerPage, pagination.Page)
	rows, err := repository.db.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := domain.Product{}
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

		products = append(products, product)
	}

	return &domain.Pagination{
		Items: products,
		Total: int32(len(products)),
	}, nil
}
