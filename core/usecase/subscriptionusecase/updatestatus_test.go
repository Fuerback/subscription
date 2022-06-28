package subscriptionusecase

import (
	"fmt"
	"testing"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/domain/mocks"
	"github.com/Fuerback/subscription/core/dto"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUpdateStatus(t *testing.T) {
	fakeId := uuid.NewString()
	status := &dto.UpdateSubscriptionStatus{Status: domain.Active}
	subscription := &domain.SubscriptionDetails{Status: domain.Paused}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSubscriptionRepository := mocks.NewMockSubscriptionRepository(mockCtrl)
	mockSubscriptionRepository.EXPECT().FetchOne(fakeId).Return(subscription, nil)
	mockSubscriptionRepository.EXPECT().UpdateStatus(fakeId, status).Return(nil)

	sut := New(mockSubscriptionRepository)
	err := sut.UpdateStatus(fakeId, status)

	require.Nil(t, err)
}

func TestUpdateStatus_CancelledStatusError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeId := uuid.NewString()
	status := &dto.UpdateSubscriptionStatus{Status: domain.Active}
	subscription := &domain.SubscriptionDetails{Status: domain.Cancelled}

	mockSubscriptionRepository := mocks.NewMockSubscriptionRepository(mockCtrl)
	mockSubscriptionRepository.EXPECT().FetchOne(fakeId).Return(subscription, nil)
	mockSubscriptionRepository.EXPECT().UpdateStatus(fakeId, status).Return(nil).Times(0)

	sut := New(mockSubscriptionRepository)
	err := sut.UpdateStatus(fakeId, status)

	require.NotNil(t, err)
}

func TestUpdateStatus_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fakeId := uuid.NewString()
	status := &dto.UpdateSubscriptionStatus{Status: domain.Active}
	subscription := &domain.SubscriptionDetails{Status: domain.Paused}

	mockSubscriptionRepository := mocks.NewMockSubscriptionRepository(mockCtrl)
	mockSubscriptionRepository.EXPECT().FetchOne(fakeId).Return(subscription, nil)
	mockSubscriptionRepository.EXPECT().UpdateStatus(fakeId, status).Return(fmt.Errorf("ANY ERROR"))

	sut := New(mockSubscriptionRepository)
	err := sut.UpdateStatus(fakeId, status)

	require.NotNil(t, err)
}
