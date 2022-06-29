package subscriptionrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
)

func (repository repository) UpdateStatus(id string, status *dto.UpdateSubscriptionStatus) error {
	ctx := context.Background()

	tx, err := repository.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query, args := getQueryAndArgs(id, status)

	stmt, err := tx.PrepareContext(ctx, query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	fmt.Println(args)
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func getQueryAndArgs(id string, status *dto.UpdateSubscriptionStatus) (string, []interface{}) {
	args := make([]interface{}, 0)
	query := "UPDATE subscription SET status = ?, paused_at = ? WHERE id = ?"
	args = append(args, status.Status)
	args = append(args, time.Now().Format("2006-02-01"))

	if status.Status == domain.Paused {
		query = "UPDATE subscription SET status = ?, paused_at = ? WHERE id = ?"
	}
	if status.Status == domain.Cancelled {
		query = "UPDATE subscription SET status = ?, paused_at = NULL, cancelled_at = ? WHERE id = ?"
	}

	args = append(args, id)
	return query, args
}
