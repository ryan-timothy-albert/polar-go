// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
)

type WebhookSubscriptionCanceledPayloadType string

const (
	WebhookSubscriptionCanceledPayloadTypeSubscriptionCanceled WebhookSubscriptionCanceledPayloadType = "subscription.canceled"
)

func (e WebhookSubscriptionCanceledPayloadType) ToPointer() *WebhookSubscriptionCanceledPayloadType {
	return &e
}
func (e *WebhookSubscriptionCanceledPayloadType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "subscription.canceled":
		*e = WebhookSubscriptionCanceledPayloadType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebhookSubscriptionCanceledPayloadType: %v", v)
	}
}

// WebhookSubscriptionCanceledPayload - Sent when a subscription is canceled by the user.
// They might still have access until the end of the current period.
//
// **Discord & Slack support:** Full
type WebhookSubscriptionCanceledPayload struct {
	type_ WebhookSubscriptionCanceledPayloadType `const:"subscription.canceled" json:"type"`
	Data  Subscription                           `json:"data"`
}

func (w WebhookSubscriptionCanceledPayload) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebhookSubscriptionCanceledPayload) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebhookSubscriptionCanceledPayload) GetType() WebhookSubscriptionCanceledPayloadType {
	return WebhookSubscriptionCanceledPayloadTypeSubscriptionCanceled
}

func (o *WebhookSubscriptionCanceledPayload) GetData() Subscription {
	if o == nil {
		return Subscription{}
	}
	return o.Data
}
