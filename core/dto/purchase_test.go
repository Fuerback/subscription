package dto

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestFromJSONCreatePurchaseRequest(t *testing.T) {
	fakeItem := PurchaseRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	r := httptest.NewRequest(http.MethodPost, "/v1/product/purchase/"+uuid.NewString(), strings.NewReader(string(json)))

	itemRequest, err := FromJSONPurchaseProductRequest(r)

	require.Nil(t, err)
	require.Equal(t, itemRequest.Voucher, fakeItem.Voucher)
}

func TestFromJSONCreatePurchaseRequest_JSONDecodeError(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/v1/product/purchase/"+uuid.NewString(), strings.NewReader("{"))
	itemRequest, err := FromJSONPurchaseProductRequest(r)

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}
