package subscriptionservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fuerback/subscription/core/domain"
	"github.com/Fuerback/subscription/core/domain/mocks"
	"github.com/Fuerback/subscription/core/dto"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func setupUpdateStatus(t *testing.T) (string, dto.UpdateSubscriptionStatus, *gomock.Controller) {
	id := uuid.NewString()
	fakeUpdateStatus := dto.UpdateSubscriptionStatus{Status: domain.Active}

	mockCtrl := gomock.NewController(t)

	return id, fakeUpdateStatus, mockCtrl
}

func TestUpdateStatus(t *testing.T) {
	fakeId, fakeUpdateStatus, mock := setupUpdateStatus(t)
	defer mock.Finish()
	mockSubscriptionUseCase := mocks.NewMockSubscriptionUseCase(mock)
	mockSubscriptionUseCase.EXPECT().UpdateStatus(fakeId, gomock.Any()).Return(nil)

	sut := New(mockSubscriptionUseCase)

	payload, _ := json.Marshal(fakeUpdateStatus)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, "/v1/subscription/status/"+fakeId, bytes.NewBuffer(payload))
	r.Header.Set("Content-Type", "application/json")

	sut.UpdateStatus(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestUpdateStatus_Error(t *testing.T) {
	fakeId, fakeUpdateStatus, mock := setupUpdateStatus(t)
	fmt.Println(fakeId)
	defer mock.Finish()
	mockSubscriptionUseCase := mocks.NewMockSubscriptionUseCase(mock)
	mockSubscriptionUseCase.EXPECT().UpdateStatus(fakeId, gomock.Any()).Return(fmt.Errorf("ANY ERROR"))

	sut := New(mockSubscriptionUseCase)

	payload, _ := json.Marshal(fakeUpdateStatus)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, "/v1/subscription/status/"+fakeId, bytes.NewBuffer(payload))
	r.Header.Set("Content-Type", "application/json")
	sut.UpdateStatus(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.NotEqual(t, http.StatusOK, res.StatusCode)
}
