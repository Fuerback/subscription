package dto

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestFromJSONCreateProductRequest(t *testing.T) {
	fakeItem := CreateProductRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := FromJSONCreateProductRequest(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, itemRequest.Name, fakeItem.Name)
	require.Equal(t, itemRequest.Price, fakeItem.Price)
	require.Equal(t, itemRequest.Period, fakeItem.Period)
	require.Equal(t, itemRequest.Active, fakeItem.Active)
}

func TestFromJSONCreateProductRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := FromJSONCreateProductRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}
