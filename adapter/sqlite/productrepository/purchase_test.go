package productrepository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fuerback/subscription/core/domain"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestPurchase(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := New(db)

	mock.ExpectBegin()
	mock.ExpectPrepare("insert into subscription *")
	mock.ExpectExec("insert into subscription *").WillReturnResult(sqlmock.NewResult(1, 1))

	fakeSubscription := &domain.Subscription{}
	faker.FakeData(fakeSubscription)

	_, err = repo.Purchase(fakeSubscription)
	require.Error(t, sql.ErrNoRows)
}

func TestPurchase_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := New(db)

	mock.ExpectPrepare("insert into subscription *")
	mock.ExpectExec("insert into subscription *").WillReturnError(fmt.Errorf("some error"))

	fakeSubscription := &domain.Subscription{}
	faker.FakeData(fakeSubscription)

	_, err = repo.Purchase(fakeSubscription)
	require.NotNil(t, err)
}
