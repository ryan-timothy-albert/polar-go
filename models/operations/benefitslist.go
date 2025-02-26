// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"github.com/polarsource/polar-go/models/components"
)

type BenefitsListQueryParamOrganizationIDFilterType string

const (
	BenefitsListQueryParamOrganizationIDFilterTypeStr        BenefitsListQueryParamOrganizationIDFilterType = "str"
	BenefitsListQueryParamOrganizationIDFilterTypeArrayOfStr BenefitsListQueryParamOrganizationIDFilterType = "arrayOfStr"
)

// BenefitsListQueryParamOrganizationIDFilter - Filter by organization ID.
type BenefitsListQueryParamOrganizationIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type BenefitsListQueryParamOrganizationIDFilterType
}

func CreateBenefitsListQueryParamOrganizationIDFilterStr(str string) BenefitsListQueryParamOrganizationIDFilter {
	typ := BenefitsListQueryParamOrganizationIDFilterTypeStr

	return BenefitsListQueryParamOrganizationIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateBenefitsListQueryParamOrganizationIDFilterArrayOfStr(arrayOfStr []string) BenefitsListQueryParamOrganizationIDFilter {
	typ := BenefitsListQueryParamOrganizationIDFilterTypeArrayOfStr

	return BenefitsListQueryParamOrganizationIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *BenefitsListQueryParamOrganizationIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = BenefitsListQueryParamOrganizationIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = BenefitsListQueryParamOrganizationIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for BenefitsListQueryParamOrganizationIDFilter", string(data))
}

func (u BenefitsListQueryParamOrganizationIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type BenefitsListQueryParamOrganizationIDFilter: all fields are null")
}

type BenefitTypeFilterType string

const (
	BenefitTypeFilterTypeBenefitType        BenefitTypeFilterType = "BenefitType"
	BenefitTypeFilterTypeArrayOfBenefitType BenefitTypeFilterType = "arrayOfBenefitType"
)

// BenefitTypeFilter - Filter by benefit type.
type BenefitTypeFilter struct {
	BenefitType        *components.BenefitType  `queryParam:"inline"`
	ArrayOfBenefitType []components.BenefitType `queryParam:"inline"`

	Type BenefitTypeFilterType
}

func CreateBenefitTypeFilterBenefitType(benefitType components.BenefitType) BenefitTypeFilter {
	typ := BenefitTypeFilterTypeBenefitType

	return BenefitTypeFilter{
		BenefitType: &benefitType,
		Type:        typ,
	}
}

func CreateBenefitTypeFilterArrayOfBenefitType(arrayOfBenefitType []components.BenefitType) BenefitTypeFilter {
	typ := BenefitTypeFilterTypeArrayOfBenefitType

	return BenefitTypeFilter{
		ArrayOfBenefitType: arrayOfBenefitType,
		Type:               typ,
	}
}

func (u *BenefitTypeFilter) UnmarshalJSON(data []byte) error {

	var benefitType components.BenefitType = components.BenefitType("")
	if err := utils.UnmarshalJSON(data, &benefitType, "", true, true); err == nil {
		u.BenefitType = &benefitType
		u.Type = BenefitTypeFilterTypeBenefitType
		return nil
	}

	var arrayOfBenefitType []components.BenefitType = []components.BenefitType{}
	if err := utils.UnmarshalJSON(data, &arrayOfBenefitType, "", true, true); err == nil {
		u.ArrayOfBenefitType = arrayOfBenefitType
		u.Type = BenefitTypeFilterTypeArrayOfBenefitType
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for BenefitTypeFilter", string(data))
}

func (u BenefitTypeFilter) MarshalJSON() ([]byte, error) {
	if u.BenefitType != nil {
		return utils.MarshalJSON(u.BenefitType, "", true)
	}

	if u.ArrayOfBenefitType != nil {
		return utils.MarshalJSON(u.ArrayOfBenefitType, "", true)
	}

	return nil, errors.New("could not marshal union type BenefitTypeFilter: all fields are null")
}

type BenefitsListRequest struct {
	// Filter by organization ID.
	OrganizationID *BenefitsListQueryParamOrganizationIDFilter `queryParam:"style=form,explode=true,name=organization_id"`
	// Filter by benefit type.
	TypeFilter *BenefitTypeFilter `queryParam:"style=form,explode=true,name=type"`
	// Page number, defaults to 1.
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
}

func (b BenefitsListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitsListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BenefitsListRequest) GetOrganizationID() *BenefitsListQueryParamOrganizationIDFilter {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *BenefitsListRequest) GetTypeFilter() *BenefitTypeFilter {
	if o == nil {
		return nil
	}
	return o.TypeFilter
}

func (o *BenefitsListRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *BenefitsListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type BenefitsListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ListResourceBenefit *components.ListResourceBenefit

	Next func() (*BenefitsListResponse, error)
}

func (o *BenefitsListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *BenefitsListResponse) GetListResourceBenefit() *components.ListResourceBenefit {
	if o == nil {
		return nil
	}
	return o.ListResourceBenefit
}
