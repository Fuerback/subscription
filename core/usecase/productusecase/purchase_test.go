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

func TestPurchase(t *testing.T) {
	purchaseRequest := dto.PurchaseRequest{}
	faker.FakeData(&purchaseRequest)
	subscription := domain.Subscription{}
	faker.FakeData(&subscription)
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().FetchOne(purchaseRequest.ProductID).Return(&fakeDBProduct, nil)
	mockProductRepository.EXPECT().Purchase(gomock.Any()).Return(subscription.ID, nil)

	sut := New(mockProductRepository)
	sub, err := sut.Purchase(&purchaseRequest)

	require.Nil(t, err)
	require.NotEmpty(t, sub.ID)
	require.Equal(t, subscription.ID, sub.ID)
}

func TestPurchase_Error(t *testing.T) {
	purchaseRequest := dto.PurchaseRequest{}
	faker.FakeData(&purchaseRequest)
	subscription := domain.Subscription{}
	faker.FakeData(&subscription)
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().FetchOne(purchaseRequest.ProductID).Return(&fakeDBProduct, nil)
	mockProductRepository.EXPECT().Purchase(gomock.Any()).Return("", fmt.Errorf("ANY ERROR"))

	sut := New(mockProductRepository)
	product, err := sut.Purchase(&purchaseRequest)

	require.NotNil(t, err)
	require.Nil(t, product)
}
