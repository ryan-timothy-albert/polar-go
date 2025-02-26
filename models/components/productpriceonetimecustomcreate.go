// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
)

type ProductPriceOneTimeCustomCreateType string

const (
	ProductPriceOneTimeCustomCreateTypeOneTime ProductPriceOneTimeCustomCreateType = "one_time"
)

func (e ProductPriceOneTimeCustomCreateType) ToPointer() *ProductPriceOneTimeCustomCreateType {
	return &e
}
func (e *ProductPriceOneTimeCustomCreateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "one_time":
		*e = ProductPriceOneTimeCustomCreateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductPriceOneTimeCustomCreateType: %v", v)
	}
}

type ProductPriceOneTimeCustomCreateAmountType string

const (
	ProductPriceOneTimeCustomCreateAmountTypeCustom ProductPriceOneTimeCustomCreateAmountType = "custom"
)

func (e ProductPriceOneTimeCustomCreateAmountType) ToPointer() *ProductPriceOneTimeCustomCreateAmountType {
	return &e
}
func (e *ProductPriceOneTimeCustomCreateAmountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "custom":
		*e = ProductPriceOneTimeCustomCreateAmountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductPriceOneTimeCustomCreateAmountType: %v", v)
	}
}

// ProductPriceOneTimeCustomCreate - Schema to create a pay-what-you-want price for a one-time product.
type ProductPriceOneTimeCustomCreate struct {
	type_      ProductPriceOneTimeCustomCreateType       `const:"one_time" json:"type"`
	amountType ProductPriceOneTimeCustomCreateAmountType `const:"custom" json:"amount_type"`
	// The currency. Currently, only `usd` is supported.
	PriceCurrency *string `default:"usd" json:"price_currency"`
	// The minimum amount the customer can pay.
	MinimumAmount *int64 `json:"minimum_amount,omitempty"`
	// The maximum amount the customer can pay.
	MaximumAmount *int64 `json:"maximum_amount,omitempty"`
	// The initial amount shown to the customer.
	PresetAmount *int64 `json:"preset_amount,omitempty"`
}

func (p ProductPriceOneTimeCustomCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductPriceOneTimeCustomCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ProductPriceOneTimeCustomCreate) GetType() ProductPriceOneTimeCustomCreateType {
	return ProductPriceOneTimeCustomCreateTypeOneTime
}

func (o *ProductPriceOneTimeCustomCreate) GetAmountType() ProductPriceOneTimeCustomCreateAmountType {
	return ProductPriceOneTimeCustomCreateAmountTypeCustom
}

func (o *ProductPriceOneTimeCustomCreate) GetPriceCurrency() *string {
	if o == nil {
		return nil
	}
	return o.PriceCurrency
}

func (o *ProductPriceOneTimeCustomCreate) GetMinimumAmount() *int64 {
	if o == nil {
		return nil
	}
	return o.MinimumAmount
}

func (o *ProductPriceOneTimeCustomCreate) GetMaximumAmount() *int64 {
	if o == nil {
		return nil
	}
	return o.MaximumAmount
}

func (o *ProductPriceOneTimeCustomCreate) GetPresetAmount() *int64 {
	if o == nil {
		return nil
	}
	return o.PresetAmount
}
