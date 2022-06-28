package productrepository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fuerback/subscription/core/dto"
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

	mock.ExpectQuery("select(.*) from product where *").WillReturnRows(sqlmock.NewRows([]string{}))

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

	pagination := &dto.PaginationRequestParms{
		Page:    0,
		PerPage: 10,
	}

	mock.ExpectQuery("select(.*) from product where *").WillReturnError(fmt.Errorf("some error"))

	_, err = repo.Fetch(pagination)
	require.NotNil(t, err)
}
