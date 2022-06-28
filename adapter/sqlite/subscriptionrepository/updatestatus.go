package subscriptionrepository

import (
	"context"
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

	var pausedAt string
	query := "UPDATE subscription SET status = ?, paused_at = ? WHERE id = ?"

	if status.Status == domain.Paused {
		pausedAt = time.Now().Format("2006-02-01")
		query = "UPDATE subscription SET status = ?, paused_at = ? WHERE id = ?"
	}

	stmt, err := tx.PrepareContext(ctx, query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, status.Status, pausedAt, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
