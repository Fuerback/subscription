package productusecase

import (
	"fmt"
	"testing"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/domain/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestFetchOne(t *testing.T) {
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeDBProduct)
	fakeId := uuid.NewString()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().FetchOne(fakeId).Return(&fakeDBProduct, nil)

	sut := New(mockProductRepository)
	product, err := sut.FetchOne(fakeId)

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeDBProduct.Name)
	require.Equal(t, product.Price, fakeDBProduct.Price)
	require.Equal(t, product.Period, fakeDBProduct.Period)
	require.Equal(t, product.Active, fakeDBProduct.Active)
}

func TestFetchOne_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeId := uuid.NewString()

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().FetchOne(fakeId).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockProductRepository)
	product, err := sut.FetchOne(fakeId)

	require.NotNil(t, err)
	require.Nil(t, product)
}
