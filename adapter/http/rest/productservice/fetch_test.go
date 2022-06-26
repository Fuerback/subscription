package productservice

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/domain/mocks"
	"github.com/Fuerback/subscription/core/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupFetch(t *testing.T) (dto.PaginationRequestParms, domain.Product, *gomock.Controller) {
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:    1,
		PerPage: 10,
	}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)

	return fakePaginationRequestParams, fakeProduct, mockCtrl
}

func TestFetch(t *testing.T) {
	fakePaginationRequestParams, fakeProduct, mock := setupFetch(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(&domain.Pagination{
		Items: []domain.Product{fakeProduct},
		Total: 1,
	}, nil)

	sut := New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/product", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("perPage", "10")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, 200, res.StatusCode)
}

func TestFetch_PorductError(t *testing.T) {
	fakePaginationRequestParams, _, mock := setupFetch(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/product", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("perPage", "10")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.NotEqual(t, 200, res.StatusCode)
}
