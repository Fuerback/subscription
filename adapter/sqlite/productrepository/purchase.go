package productrepository

import (
	"context"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/google/uuid"
)

func (repository repository) Purchase(subscription *domain.Subscription) (string, error) {
	ctx := context.Background()

	tx, err := repository.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	stmt, err := tx.PrepareContext(ctx, "insert into subscription(id, startsAt, endsAt, status, voucher_id, account_id, product_id) values (?,?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return "", err
	}

	subscriptionId := uuid.NewString()
	_, err = stmt.ExecContext(ctx, subscriptionId, subscription.StartsAt, subscription.EndsAt, subscription.Status, subscription.Voucher, subscription.Account, subscription.Product)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()

	return subscriptionId, nil
}
