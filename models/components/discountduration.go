// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type DiscountDuration string

const (
	DiscountDurationOnce      DiscountDuration = "once"
	DiscountDurationForever   DiscountDuration = "forever"
	DiscountDurationRepeating DiscountDuration = "repeating"
)

func (e DiscountDuration) ToPointer() *DiscountDuration {
	return &e
}
func (e *DiscountDuration) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "once":
		fallthrough
	case "forever":
		fallthrough
	case "repeating":
		*e = DiscountDuration(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DiscountDuration: %v", v)
	}
}
