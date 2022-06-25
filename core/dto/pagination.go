package dto

import (
	"net/http"
	"strconv"
)

// PaginationRequestParms is an representation query string params to filter and paginate products
type PaginationRequestParms struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// FromValuePaginationRequestParams converts query string params to a PaginationRequestParms struct
func FromValuePaginationRequestParams(request *http.Request) (*PaginationRequestParms, error) {
	page, err := strconv.Atoi(request.FormValue("page"))
	if err != nil {
		page = 0
	}
	perPage, err := strconv.Atoi(request.FormValue("perPage"))
	if err != nil {
		perPage = 10
	}

	paginationRequestParms := PaginationRequestParms{
		Page:    page,
		PerPage: perPage,
	}

	return &paginationRequestParms, nil
}
