package productusecase

import (
	"fmt"
	"testing"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/domain/mocks"
	"github.com/Fuerback/subscription/core/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:    1,
		PerPage: 10,
	}
	fakeDBProduct := domain.Product{}

	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Fetch(&fakePaginationRequestParams).Return([]domain.Product{fakeDBProduct}, nil)

	sut := New(mockProductRepository)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.Nil(t, err)

	for _, product := range products {
		require.Nil(t, err)
		require.NotEmpty(t, product.ID)
		require.Equal(t, product.Name, fakeDBProduct.Name)
		require.Equal(t, product.Price, fakeDBProduct.Price)
		require.Equal(t, product.Period, fakeDBProduct.Period)
		require.Equal(t, product.Active, fakeDBProduct.Active)
	}
}

func TestFetch_Error(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:    1,
		PerPage: 10,
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockProductRepository)
	product, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)
	require.Nil(t, product)
}
