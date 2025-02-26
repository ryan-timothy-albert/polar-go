// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"fmt"
)

type AlreadyCanceledSubscriptionError string

const (
	AlreadyCanceledSubscriptionErrorAlreadyCanceledSubscription AlreadyCanceledSubscriptionError = "AlreadyCanceledSubscription"
)

func (e AlreadyCanceledSubscriptionError) ToPointer() *AlreadyCanceledSubscriptionError {
	return &e
}
func (e *AlreadyCanceledSubscriptionError) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AlreadyCanceledSubscription":
		*e = AlreadyCanceledSubscriptionError(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AlreadyCanceledSubscriptionError: %v", v)
	}
}

type AlreadyCanceledSubscription struct {
	Error_ AlreadyCanceledSubscriptionError `const:"AlreadyCanceledSubscription" json:"error"`
	Detail string                           `json:"detail"`
}

var _ error = &AlreadyCanceledSubscription{}

func (e *AlreadyCanceledSubscription) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
