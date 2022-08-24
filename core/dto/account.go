package dto

import "github.com/Fuerback/subscription/core/domain"

type AccountResponse struct {
	Name string `json:"name"`
}

func FromDomainToDtoAccount(account domain.Account) AccountResponse {
	return AccountResponse{
		Name: account.Name,
	}
}
