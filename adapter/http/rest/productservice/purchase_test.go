package productservice

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
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func setupPurchase(t *testing.T) (string, domain.Subscription, dto.PurchaseRequest, *gomock.Controller) {
	id := uuid.NewString()
	fakeSubscription := domain.Subscription{}
	faker.FakeData(&fakeSubscription)
	purchaseRequest := dto.PurchaseRequest{}
	faker.FakeData(purchaseRequest)

	mockCtrl := gomock.NewController(t)

	return id, fakeSubscription, purchaseRequest, mockCtrl
}

func TestPurchase(t *testing.T) {
	fakeId, fakeSubscription, fakePurchase, mock := setupPurchase(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Purchase(gomock.Any()).Return(&fakeSubscription, nil)

	sut := New(mockProductUseCase)

	payload, _ := json.Marshal(fakePurchase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/product/purchase/"+fakeId, bytes.NewBuffer(payload))
	r.Header.Set("x-account-id", uuid.NewString())
	r.Header.Set("Content-Type", "application/json")

	sut.Purchase(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, 200, res.StatusCode)
}

func TestPurchase_PorductError(t *testing.T) {
	fakeId, _, fakePurchase, mock := setupPurchase(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Purchase(gomock.Any()).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockProductUseCase)

	payload, _ := json.Marshal(fakePurchase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/product/purchase/"+fakeId, bytes.NewBuffer(payload))
	r.Header.Set("x-account-id", uuid.NewString())
	r.Header.Set("Content-Type", "application/json")
	sut.Purchase(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.NotEqual(t, 200, res.StatusCode)
}
