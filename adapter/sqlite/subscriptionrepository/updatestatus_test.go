package subscriptionrepository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUpdateStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := New(db)
	id := uuid.NewString()
	status := dto.UpdateSubscriptionStatus{Status: domain.Active}

	mock.ExpectQuery("UPDATE subscription SET status *").WillReturnRows(sqlmock.NewRows([]string{}))

	err = repo.UpdateStatus(id, &status)
	require.Error(t, sql.ErrNoRows)
}

func TestUpdateStatus_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := New(db)
	id := uuid.NewString()
	status := dto.UpdateSubscriptionStatus{Status: domain.Active}

	mock.ExpectQuery("UPDATE subscription SET status *").WillReturnError(fmt.Errorf("some error"))

	err = repo.UpdateStatus(id, &status)
	require.NotNil(t, err)
}
