package dto

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromValuePaginationRequestParams(t *testing.T) {
	fakeRequest := httptest.NewRequest(http.MethodGet, "/product", nil)
	queryStringParams := fakeRequest.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("perPage", "10")
	fakeRequest.URL.RawQuery = queryStringParams.Encode()

	paginationRequest, err := FromValuePaginationRequestParams(fakeRequest)

	require.Nil(t, err)
	require.Equal(t, paginationRequest.Page, 1)
	require.Equal(t, paginationRequest.PerPage, 10)
}
