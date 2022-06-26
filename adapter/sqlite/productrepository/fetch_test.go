package productrepository

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fuerback/subscription/core/dto"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
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

	mock.ExpectQuery("select(.*) from product limit *").WillReturnRows(sqlmock.NewRows([]string{}))

	_, err = repo.Fetch(pagination)
	require.Nil(t, err)

	err = mock.ExpectationsWereMet()
	require.Nil(t, err)
}

func TestFetch_Error(t *testing.T) {
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

	mock.ExpectQuery("select(.*) from product limit *").WillReturnError(fmt.Errorf("some error"))

	_, err = repo.Fetch(pagination)
	require.NotNil(t, err)

	err = mock.ExpectationsWereMet()
	require.Nil(t, err)
}
