package dto

import (
	"encoding/json"
	"io"
)

// UpdateSubscriptionStatus is an representation request body to update subscription status
type UpdateSubscriptionStatus struct {
	Status string `json:"status" validate:"required,oneof=PAUSED ACTIVE"`
}

// FromJSONUpdateSubscriptionStatusRequest converts json body request to a UpdateSubscriptionStatus struct
func FromJSONUpdateSubscriptionStatusRequest(body io.ReadCloser) (*UpdateSubscriptionStatus, error) {
	updatestatusRequest := UpdateSubscriptionStatus{}
	if err := json.NewDecoder(body).Decode(&updatestatusRequest); err != nil {
		return nil, err
	}

	return &updatestatusRequest, nil
}
