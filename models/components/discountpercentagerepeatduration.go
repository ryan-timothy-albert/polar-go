// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"time"
)

type DiscountPercentageRepeatDurationMetadataType string

const (
	DiscountPercentageRepeatDurationMetadataTypeStr     DiscountPercentageRepeatDurationMetadataType = "str"
	DiscountPercentageRepeatDurationMetadataTypeInteger DiscountPercentageRepeatDurationMetadataType = "integer"
	DiscountPercentageRepeatDurationMetadataTypeBoolean DiscountPercentageRepeatDurationMetadataType = "boolean"
)

type DiscountPercentageRepeatDurationMetadata struct {
	Str     *string `queryParam:"inline"`
	Integer *int64  `queryParam:"inline"`
	Boolean *bool   `queryParam:"inline"`

	Type DiscountPercentageRepeatDurationMetadataType
}

func CreateDiscountPercentageRepeatDurationMetadataStr(str string) DiscountPercentageRepeatDurationMetadata {
	typ := DiscountPercentageRepeatDurationMetadataTypeStr

	return DiscountPercentageRepeatDurationMetadata{
		Str:  &str,
		Type: typ,
	}
}

func CreateDiscountPercentageRepeatDurationMetadataInteger(integer int64) DiscountPercentageRepeatDurationMetadata {
	typ := DiscountPercentageRepeatDurationMetadataTypeInteger

	return DiscountPercentageRepeatDurationMetadata{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateDiscountPercentageRepeatDurationMetadataBoolean(boolean bool) DiscountPercentageRepeatDurationMetadata {
	typ := DiscountPercentageRepeatDurationMetadataTypeBoolean

	return DiscountPercentageRepeatDurationMetadata{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *DiscountPercentageRepeatDurationMetadata) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = DiscountPercentageRepeatDurationMetadataTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = DiscountPercentageRepeatDurationMetadataTypeInteger
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = DiscountPercentageRepeatDurationMetadataTypeBoolean
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DiscountPercentageRepeatDurationMetadata", string(data))
}

func (u DiscountPercentageRepeatDurationMetadata) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type DiscountPercentageRepeatDurationMetadata: all fields are null")
}

// DiscountPercentageRepeatDuration - Schema for a percentage discount that is applied on every invoice
// for a certain number of months.
type DiscountPercentageRepeatDuration struct {
	Duration         DiscountDuration `json:"duration"`
	DurationInMonths int64            `json:"duration_in_months"`
	Type             DiscountType     `json:"type"`
	BasisPoints      int64            `json:"basis_points"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at"`
	// Last modification timestamp of the object.
	ModifiedAt *time.Time `json:"modified_at"`
	// The ID of the object.
	ID       string                                              `json:"id"`
	Metadata map[string]DiscountPercentageRepeatDurationMetadata `json:"metadata"`
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
	OrganizationID string            `json:"organization_id"`
	Products       []DiscountProduct `json:"products"`
}

func (d DiscountPercentageRepeatDuration) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DiscountPercentageRepeatDuration) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DiscountPercentageRepeatDuration) GetDuration() DiscountDuration {
	if o == nil {
		return DiscountDuration("")
	}
	return o.Duration
}

func (o *DiscountPercentageRepeatDuration) GetDurationInMonths() int64 {
	if o == nil {
		return 0
	}
	return o.DurationInMonths
}

func (o *DiscountPercentageRepeatDuration) GetType() DiscountType {
	if o == nil {
		return DiscountType("")
	}
	return o.Type
}

func (o *DiscountPercentageRepeatDuration) GetBasisPoints() int64 {
	if o == nil {
		return 0
	}
	return o.BasisPoints
}

func (o *DiscountPercentageRepeatDuration) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *DiscountPercentageRepeatDuration) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *DiscountPercentageRepeatDuration) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DiscountPercentageRepeatDuration) GetMetadata() map[string]DiscountPercentageRepeatDurationMetadata {
	if o == nil {
		return map[string]DiscountPercentageRepeatDurationMetadata{}
	}
	return o.Metadata
}

func (o *DiscountPercentageRepeatDuration) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DiscountPercentageRepeatDuration) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *DiscountPercentageRepeatDuration) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *DiscountPercentageRepeatDuration) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *DiscountPercentageRepeatDuration) GetMaxRedemptions() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRedemptions
}

func (o *DiscountPercentageRepeatDuration) GetRedemptionsCount() int64 {
	if o == nil {
		return 0
	}
	return o.RedemptionsCount
}

func (o *DiscountPercentageRepeatDuration) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *DiscountPercentageRepeatDuration) GetProducts() []DiscountProduct {
	if o == nil {
		return []DiscountProduct{}
	}
	return o.Products
}
