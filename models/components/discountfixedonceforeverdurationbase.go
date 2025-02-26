// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"time"
)

type DiscountFixedOnceForeverDurationBaseMetadataType string

const (
	DiscountFixedOnceForeverDurationBaseMetadataTypeStr     DiscountFixedOnceForeverDurationBaseMetadataType = "str"
	DiscountFixedOnceForeverDurationBaseMetadataTypeInteger DiscountFixedOnceForeverDurationBaseMetadataType = "integer"
	DiscountFixedOnceForeverDurationBaseMetadataTypeBoolean DiscountFixedOnceForeverDurationBaseMetadataType = "boolean"
)

type DiscountFixedOnceForeverDurationBaseMetadata struct {
	Str     *string `queryParam:"inline"`
	Integer *int64  `queryParam:"inline"`
	Boolean *bool   `queryParam:"inline"`

	Type DiscountFixedOnceForeverDurationBaseMetadataType
}

func CreateDiscountFixedOnceForeverDurationBaseMetadataStr(str string) DiscountFixedOnceForeverDurationBaseMetadata {
	typ := DiscountFixedOnceForeverDurationBaseMetadataTypeStr

	return DiscountFixedOnceForeverDurationBaseMetadata{
		Str:  &str,
		Type: typ,
	}
}

func CreateDiscountFixedOnceForeverDurationBaseMetadataInteger(integer int64) DiscountFixedOnceForeverDurationBaseMetadata {
	typ := DiscountFixedOnceForeverDurationBaseMetadataTypeInteger

	return DiscountFixedOnceForeverDurationBaseMetadata{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateDiscountFixedOnceForeverDurationBaseMetadataBoolean(boolean bool) DiscountFixedOnceForeverDurationBaseMetadata {
	typ := DiscountFixedOnceForeverDurationBaseMetadataTypeBoolean

	return DiscountFixedOnceForeverDurationBaseMetadata{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *DiscountFixedOnceForeverDurationBaseMetadata) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = DiscountFixedOnceForeverDurationBaseMetadataTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = DiscountFixedOnceForeverDurationBaseMetadataTypeInteger
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = DiscountFixedOnceForeverDurationBaseMetadataTypeBoolean
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DiscountFixedOnceForeverDurationBaseMetadata", string(data))
}

func (u DiscountFixedOnceForeverDurationBaseMetadata) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type DiscountFixedOnceForeverDurationBaseMetadata: all fields are null")
}

type DiscountFixedOnceForeverDurationBase struct {
	Duration DiscountDuration `json:"duration"`
	Type     DiscountType     `json:"type"`
	Amount   int64            `json:"amount"`
	Currency string           `json:"currency"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at"`
	// Last modification timestamp of the object.
	ModifiedAt *time.Time `json:"modified_at"`
	// The ID of the object.
	ID       string                                                  `json:"id"`
	Metadata map[string]DiscountFixedOnceForeverDurationBaseMetadata `json:"metadata"`
	// Name of the discount. Will be displayed to the customer when the discount is applied.
	Name string `json:"name"`
	// Code customers can use to apply the discount during checkout.
	Code *string `json:"code"`
	// Timestamp after which the discount is redeemable.
	StartsAt *time.Time `json:"starts_at"`
	// Timestamp after which the discount is no longer redeemable.
	EndsAt *time.Time `json:"ends_at"`
	// Maximum number of times the discount can be redeemed.
	MaxRedemptions *int64 `json:"max_redemptions"`
	// Number of times the discount has been redeemed.
	RedemptionsCount int64 `json:"redemptions_count"`
	// The organization ID.
	OrganizationID string `json:"organization_id"`
}

func (d DiscountFixedOnceForeverDurationBase) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DiscountFixedOnceForeverDurationBase) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DiscountFixedOnceForeverDurationBase) GetDuration() DiscountDuration {
	if o == nil {
		return DiscountDuration("")
	}
	return o.Duration
}

func (o *DiscountFixedOnceForeverDurationBase) GetType() DiscountType {
	if o == nil {
		return DiscountType("")
	}
	return o.Type
}

func (o *DiscountFixedOnceForeverDurationBase) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

func (o *DiscountFixedOnceForeverDurationBase) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *DiscountFixedOnceForeverDurationBase) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *DiscountFixedOnceForeverDurationBase) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *DiscountFixedOnceForeverDurationBase) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DiscountFixedOnceForeverDurationBase) GetMetadata() map[string]DiscountFixedOnceForeverDurationBaseMetadata {
	if o == nil {
		return map[string]DiscountFixedOnceForeverDurationBaseMetadata{}
	}
	return o.Metadata
}

func (o *DiscountFixedOnceForeverDurationBase) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DiscountFixedOnceForeverDurationBase) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *DiscountFixedOnceForeverDurationBase) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *DiscountFixedOnceForeverDurationBase) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *DiscountFixedOnceForeverDurationBase) GetMaxRedemptions() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRedemptions
}

func (o *DiscountFixedOnceForeverDurationBase) GetRedemptionsCount() int64 {
	if o == nil {
		return 0
	}
	return o.RedemptionsCount
}

func (o *DiscountFixedOnceForeverDurationBase) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}
