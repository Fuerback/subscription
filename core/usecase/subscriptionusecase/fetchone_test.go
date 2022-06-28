package subscriptionusecase

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
	fakeDBSubscription := domain.SubscriptionDetails{}
	faker.FakeData(&fakeDBSubscription)
	fakeId := uuid.NewString()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSubscriptionRepository := mocks.NewMockSubscriptionRepository(mockCtrl)
	mockSubscriptionRepository.EXPECT().FetchOne(fakeId).Return(&fakeDBSubscription, nil)

	sut := New(mockSubscriptionRepository)
	subscription, err := sut.FetchOne(fakeId)

	require.Nil(t, err)
	require.NotEmpty(t, subscription.ID)
	require.Equal(t, subscription.StartsAt, fakeDBSubscription.StartsAt)
	require.Equal(t, subscription.EndsAt, fakeDBSubscription.EndsAt)
	require.Equal(t, subscription.Status, fakeDBSubscription.Status)
	require.Equal(t, subscription.Voucher, fakeDBSubscription.Voucher)
}

func TestFetchOne_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeId := uuid.NewString()

	mockSubscriptionRepository := mocks.NewMockSubscriptionRepository(mockCtrl)
	mockSubscriptionRepository.EXPECT().FetchOne(fakeId).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockSubscriptionRepository)
	product, err := sut.FetchOne(fakeId)

	require.NotNil(t, err)
	require.Nil(t, product)
}
