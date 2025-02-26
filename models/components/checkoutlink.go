// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"time"
)

type CheckoutLinkMetadataType string

const (
	CheckoutLinkMetadataTypeStr     CheckoutLinkMetadataType = "str"
	CheckoutLinkMetadataTypeInteger CheckoutLinkMetadataType = "integer"
	CheckoutLinkMetadataTypeBoolean CheckoutLinkMetadataType = "boolean"
)

type CheckoutLinkMetadata struct {
	Str     *string `queryParam:"inline"`
	Integer *int64  `queryParam:"inline"`
	Boolean *bool   `queryParam:"inline"`

	Type CheckoutLinkMetadataType
}

func CreateCheckoutLinkMetadataStr(str string) CheckoutLinkMetadata {
	typ := CheckoutLinkMetadataTypeStr

	return CheckoutLinkMetadata{
		Str:  &str,
		Type: typ,
	}
}

func CreateCheckoutLinkMetadataInteger(integer int64) CheckoutLinkMetadata {
	typ := CheckoutLinkMetadataTypeInteger

	return CheckoutLinkMetadata{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateCheckoutLinkMetadataBoolean(boolean bool) CheckoutLinkMetadata {
	typ := CheckoutLinkMetadataTypeBoolean

	return CheckoutLinkMetadata{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *CheckoutLinkMetadata) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CheckoutLinkMetadataTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = CheckoutLinkMetadataTypeInteger
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = CheckoutLinkMetadataTypeBoolean
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CheckoutLinkMetadata", string(data))
}

func (u CheckoutLinkMetadata) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type CheckoutLinkMetadata: all fields are null")
}

type CheckoutLinkDiscountType string

const (
	CheckoutLinkDiscountTypeDiscountFixedOnceForeverDurationBase      CheckoutLinkDiscountType = "DiscountFixedOnceForeverDurationBase"
	CheckoutLinkDiscountTypeDiscountFixedRepeatDurationBase           CheckoutLinkDiscountType = "DiscountFixedRepeatDurationBase"
	CheckoutLinkDiscountTypeDiscountPercentageOnceForeverDurationBase CheckoutLinkDiscountType = "DiscountPercentageOnceForeverDurationBase"
	CheckoutLinkDiscountTypeDiscountPercentageRepeatDurationBase      CheckoutLinkDiscountType = "DiscountPercentageRepeatDurationBase"
)

type CheckoutLinkDiscount struct {
	DiscountFixedOnceForeverDurationBase      *DiscountFixedOnceForeverDurationBase      `queryParam:"inline"`
	DiscountFixedRepeatDurationBase           *DiscountFixedRepeatDurationBase           `queryParam:"inline"`
	DiscountPercentageOnceForeverDurationBase *DiscountPercentageOnceForeverDurationBase `queryParam:"inline"`
	DiscountPercentageRepeatDurationBase      *DiscountPercentageRepeatDurationBase      `queryParam:"inline"`

	Type CheckoutLinkDiscountType
}

func CreateCheckoutLinkDiscountDiscountFixedOnceForeverDurationBase(discountFixedOnceForeverDurationBase DiscountFixedOnceForeverDurationBase) CheckoutLinkDiscount {
	typ := CheckoutLinkDiscountTypeDiscountFixedOnceForeverDurationBase

	return CheckoutLinkDiscount{
		DiscountFixedOnceForeverDurationBase: &discountFixedOnceForeverDurationBase,
		Type:                                 typ,
	}
}

func CreateCheckoutLinkDiscountDiscountFixedRepeatDurationBase(discountFixedRepeatDurationBase DiscountFixedRepeatDurationBase) CheckoutLinkDiscount {
	typ := CheckoutLinkDiscountTypeDiscountFixedRepeatDurationBase

	return CheckoutLinkDiscount{
		DiscountFixedRepeatDurationBase: &discountFixedRepeatDurationBase,
		Type:                            typ,
	}
}

func CreateCheckoutLinkDiscountDiscountPercentageOnceForeverDurationBase(discountPercentageOnceForeverDurationBase DiscountPercentageOnceForeverDurationBase) CheckoutLinkDiscount {
	typ := CheckoutLinkDiscountTypeDiscountPercentageOnceForeverDurationBase

	return CheckoutLinkDiscount{
		DiscountPercentageOnceForeverDurationBase: &discountPercentageOnceForeverDurationBase,
		Type: typ,
	}
}

func CreateCheckoutLinkDiscountDiscountPercentageRepeatDurationBase(discountPercentageRepeatDurationBase DiscountPercentageRepeatDurationBase) CheckoutLinkDiscount {
	typ := CheckoutLinkDiscountTypeDiscountPercentageRepeatDurationBase

	return CheckoutLinkDiscount{
		DiscountPercentageRepeatDurationBase: &discountPercentageRepeatDurationBase,
		Type:                                 typ,
	}
}

func (u *CheckoutLinkDiscount) UnmarshalJSON(data []byte) error {

	var discountPercentageOnceForeverDurationBase DiscountPercentageOnceForeverDurationBase = DiscountPercentageOnceForeverDurationBase{}
	if err := utils.UnmarshalJSON(data, &discountPercentageOnceForeverDurationBase, "", true, true); err == nil {
		u.DiscountPercentageOnceForeverDurationBase = &discountPercentageOnceForeverDurationBase
		u.Type = CheckoutLinkDiscountTypeDiscountPercentageOnceForeverDurationBase
		return nil
	}

	var discountFixedOnceForeverDurationBase DiscountFixedOnceForeverDurationBase = DiscountFixedOnceForeverDurationBase{}
	if err := utils.UnmarshalJSON(data, &discountFixedOnceForeverDurationBase, "", true, true); err == nil {
		u.DiscountFixedOnceForeverDurationBase = &discountFixedOnceForeverDurationBase
		u.Type = CheckoutLinkDiscountTypeDiscountFixedOnceForeverDurationBase
		return nil
	}

	var discountPercentageRepeatDurationBase DiscountPercentageRepeatDurationBase = DiscountPercentageRepeatDurationBase{}
	if err := utils.UnmarshalJSON(data, &discountPercentageRepeatDurationBase, "", true, true); err == nil {
		u.DiscountPercentageRepeatDurationBase = &discountPercentageRepeatDurationBase
		u.Type = CheckoutLinkDiscountTypeDiscountPercentageRepeatDurationBase
		return nil
	}

	var discountFixedRepeatDurationBase DiscountFixedRepeatDurationBase = DiscountFixedRepeatDurationBase{}
	if err := utils.UnmarshalJSON(data, &discountFixedRepeatDurationBase, "", true, true); err == nil {
		u.DiscountFixedRepeatDurationBase = &discountFixedRepeatDurationBase
		u.Type = CheckoutLinkDiscountTypeDiscountFixedRepeatDurationBase
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CheckoutLinkDiscount", string(data))
}

func (u CheckoutLinkDiscount) MarshalJSON() ([]byte, error) {
	if u.DiscountFixedOnceForeverDurationBase != nil {
		return utils.MarshalJSON(u.DiscountFixedOnceForeverDurationBase, "", true)
	}

	if u.DiscountFixedRepeatDurationBase != nil {
		return utils.MarshalJSON(u.DiscountFixedRepeatDurationBase, "", true)
	}

	if u.DiscountPercentageOnceForeverDurationBase != nil {
		return utils.MarshalJSON(u.DiscountPercentageOnceForeverDurationBase, "", true)
	}

	if u.DiscountPercentageRepeatDurationBase != nil {
		return utils.MarshalJSON(u.DiscountPercentageRepeatDurationBase, "", true)
	}

	return nil, errors.New("could not marshal union type CheckoutLinkDiscount: all fields are null")
}

// CheckoutLink - Checkout link data.
type CheckoutLink struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at"`
	// Last modification timestamp of the object.
	ModifiedAt *time.Time `json:"modified_at"`
	// The ID of the object.
	ID               string                          `json:"id"`
	Metadata         map[string]CheckoutLinkMetadata `json:"metadata"`
	paymentProcessor PaymentProcessor                `const:"stripe" json:"payment_processor"`
	// Client secret used to access the checkout link.
	ClientSecret string `json:"client_secret"`
	// URL where the customer will be redirected after a successful payment.
	SuccessURL *string `json:"success_url"`
	// Optional label to distinguish links internally
	Label *string `json:"label"`
	// Whether to allow the customer to apply discount codes. If you apply a discount through `discount_id`, it'll still be applied, but the customer won't be able to change it.
	AllowDiscountCodes bool `json:"allow_discount_codes"`
	// ID of the product to checkout.
	ProductID string `json:"product_id"`
	// ID of the product price to checkout. First available price will be selected unless an explicit price ID is set.
	ProductPriceID *string `json:"product_price_id"`
	// ID of the discount to apply to the checkout. If the discount is not applicable anymore when opening the checkout link, it'll be ignored.
	DiscountID *string `json:"discount_id"`
	// Product data for a checkout link.
	Product      CheckoutLinkProduct   `json:"product"`
	ProductPrice *ProductPrice         `json:"product_price"`
	Discount     *CheckoutLinkDiscount `json:"discount"`
	URL          string                `json:"url"`
}

func (c CheckoutLink) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CheckoutLink) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CheckoutLink) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CheckoutLink) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *CheckoutLink) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CheckoutLink) GetMetadata() map[string]CheckoutLinkMetadata {
	if o == nil {
		return map[string]CheckoutLinkMetadata{}
	}
	return o.Metadata
}

func (o *CheckoutLink) GetPaymentProcessor() PaymentProcessor {
	return PaymentProcessorStripe
}

func (o *CheckoutLink) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *CheckoutLink) GetSuccessURL() *string {
	if o == nil {
		return nil
	}
	return o.SuccessURL
}

func (o *CheckoutLink) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *CheckoutLink) GetAllowDiscountCodes() bool {
	if o == nil {
		return false
	}
	return o.AllowDiscountCodes
}

func (o *CheckoutLink) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

func (o *CheckoutLink) GetProductPriceID() *string {
	if o == nil {
		return nil
	}
	return o.ProductPriceID
}

func (o *CheckoutLink) GetDiscountID() *string {
	if o == nil {
		return nil
	}
	return o.DiscountID
}

func (o *CheckoutLink) GetProduct() CheckoutLinkProduct {
	if o == nil {
		return CheckoutLinkProduct{}
	}
	return o.Product
}

func (o *CheckoutLink) GetProductPrice() *ProductPrice {
	if o == nil {
		return nil
	}
	return o.ProductPrice
}

func (o *CheckoutLink) GetDiscount() *CheckoutLinkDiscount {
	if o == nil {
		return nil
	}
	return o.Discount
}

func (o *CheckoutLink) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
