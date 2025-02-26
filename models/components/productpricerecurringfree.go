// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"time"
)

type ProductPriceRecurringFreeAmountType string

const (
	ProductPriceRecurringFreeAmountTypeFree ProductPriceRecurringFreeAmountType = "free"
)

func (e ProductPriceRecurringFreeAmountType) ToPointer() *ProductPriceRecurringFreeAmountType {
	return &e
}
func (e *ProductPriceRecurringFreeAmountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "free":
		*e = ProductPriceRecurringFreeAmountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductPriceRecurringFreeAmountType: %v", v)
	}
}

// ProductPriceRecurringFreeType - The type of the price.
type ProductPriceRecurringFreeType string

const (
	ProductPriceRecurringFreeTypeRecurring ProductPriceRecurringFreeType = "recurring"
)

func (e ProductPriceRecurringFreeType) ToPointer() *ProductPriceRecurringFreeType {
	return &e
}
func (e *ProductPriceRecurringFreeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "recurring":
		*e = ProductPriceRecurringFreeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductPriceRecurringFreeType: %v", v)
	}
}

// ProductPriceRecurringFree - A free recurring price for a product, i.e. a subscription.
type ProductPriceRecurringFree struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at"`
	// Last modification timestamp of the object.
	ModifiedAt *time.Time `json:"modified_at"`
	// The ID of the price.
	ID         string                              `json:"id"`
	amountType ProductPriceRecurringFreeAmountType `const:"free" json:"amount_type"`
	// Whether the price is archived and no longer available.
	IsArchived bool `json:"is_archived"`
	// The ID of the product owning the price.
	ProductID string `json:"product_id"`
	// The type of the price.
	type_             ProductPriceRecurringFreeType `const:"recurring" json:"type"`
	RecurringInterval SubscriptionRecurringInterval `json:"recurring_interval"`
}

func (p ProductPriceRecurringFree) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductPriceRecurringFree) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ProductPriceRecurringFree) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ProductPriceRecurringFree) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *ProductPriceRecurringFree) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ProductPriceRecurringFree) GetAmountType() ProductPriceRecurringFreeAmountType {
	return ProductPriceRecurringFreeAmountTypeFree
}

func (o *ProductPriceRecurringFree) GetIsArchived() bool {
	if o == nil {
		return false
	}
	return o.IsArchived
}

func (o *ProductPriceRecurringFree) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

func (o *ProductPriceRecurringFree) GetType() ProductPriceRecurringFreeType {
	return ProductPriceRecurringFreeTypeRecurring
}

func (o *ProductPriceRecurringFree) GetRecurringInterval() SubscriptionRecurringInterval {
	if o == nil {
		return SubscriptionRecurringInterval("")
	}
	return o.RecurringInterval
}
