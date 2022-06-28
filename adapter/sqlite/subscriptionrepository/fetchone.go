package subscriptionrepository

import (
	"context"

	"github.com/Fuerback/subscription/core/domain"
)

const (
	GetSubscriptionDetails = `
	select s.id, cast(s.startsAt as text), cast(s.endsAt as text), s.status, s.voucher_id, p.name, p.period, p.price, a.name 
	from subscription s 
	join product as p on p.id = s.product_id 
	join account as a on a.id = s.account_id 
	where s.id = ?
	`
)

func (repository repository) FetchOne(id string) (*domain.SubscriptionDetails, error) {
	ctx := context.Background()
	subscription := &domain.SubscriptionDetails{}

	rows := repository.db.QueryRowContext(
		ctx,
		GetSubscriptionDetails,
		id,
	)
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	err := rows.Scan(
		&subscription.ID,
		&subscription.StartsAt,
		&subscription.EndsAt,
		&subscription.Status,
		&subscription.Voucher,
		&subscription.Product.Name,
		&subscription.Product.Period,
		&subscription.Product.Price,
		&subscription.Account.Name,
	)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
