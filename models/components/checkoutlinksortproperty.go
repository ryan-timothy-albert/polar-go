// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type CheckoutLinkSortProperty string

const (
	CheckoutLinkSortPropertyCreatedAt      CheckoutLinkSortProperty = "created_at"
	CheckoutLinkSortPropertyMinusCreatedAt CheckoutLinkSortProperty = "-created_at"
)

func (e CheckoutLinkSortProperty) ToPointer() *CheckoutLinkSortProperty {
	return &e
}
func (e *CheckoutLinkSortProperty) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		fallthrough
	case "-created_at":
		*e = CheckoutLinkSortProperty(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CheckoutLinkSortProperty: %v", v)
	}
}
