package productservice

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/domain/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func setupFetchOne(t *testing.T) (string, domain.Product, *gomock.Controller) {
	id := uuid.NewString()
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)

	return id, fakeProduct, mockCtrl
}

func TestFetchOne(t *testing.T) {
	fakeId, fakeProduct, mock := setupFetchOne(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().FetchOne(fakeId).Return(&fakeProduct, nil)

	sut := New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/product/"+fakeId, nil)
	r.Header.Set("Content-Type", "application/json")

	sut.FetchOne(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestFetchOne_PorductError(t *testing.T) {
	fakeId, _, mock := setupFetchOne(t)
	fmt.Println(fakeId)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().FetchOne(fakeId).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/product/"+fakeId, nil)
	r.Header.Set("Content-Type", "application/json")
	sut.FetchOne(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.NotEqual(t, http.StatusOK, res.StatusCode)
}
