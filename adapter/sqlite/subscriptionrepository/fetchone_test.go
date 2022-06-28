package subscriptionrepository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestFetchOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := New(db)
	id := uuid.NewString()

	mock.ExpectQuery("select(.*) from subscription *").WillReturnRows(sqlmock.NewRows([]string{}))

	_, err = repo.FetchOne(id)
	require.Error(t, sql.ErrNoRows)
}

func TestFetchOne_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := New(db)
	id := uuid.NewString()

	mock.ExpectQuery("select(.*) from subscription *").WillReturnError(fmt.Errorf("some error"))

	_, err = repo.FetchOne(id)
	require.NotNil(t, err)
}
