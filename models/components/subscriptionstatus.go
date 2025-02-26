// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type SubscriptionStatus string

const (
	SubscriptionStatusIncomplete        SubscriptionStatus = "incomplete"
	SubscriptionStatusIncompleteExpired SubscriptionStatus = "incomplete_expired"
	SubscriptionStatusTrialing          SubscriptionStatus = "trialing"
	SubscriptionStatusActive            SubscriptionStatus = "active"
	SubscriptionStatusPastDue           SubscriptionStatus = "past_due"
	SubscriptionStatusCanceled          SubscriptionStatus = "canceled"
	SubscriptionStatusUnpaid            SubscriptionStatus = "unpaid"
)

func (e SubscriptionStatus) ToPointer() *SubscriptionStatus {
	return &e
}
func (e *SubscriptionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "incomplete":
		fallthrough
	case "incomplete_expired":
		fallthrough
	case "trialing":
		fallthrough
	case "active":
		fallthrough
	case "past_due":
		fallthrough
	case "canceled":
		fallthrough
	case "unpaid":
		*e = SubscriptionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SubscriptionStatus: %v", v)
	}
}
