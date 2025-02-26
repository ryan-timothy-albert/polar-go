// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"github.com/polarsource/polar-go/models/components"
)

type CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterType string

const (
	CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterTypeStr        CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterType = "str"
	CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterTypeArrayOfStr CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterType = "arrayOfStr"
)

// CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter - Filter by organization ID.
type CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterType
}

func CreateCustomerPortalLicenseKeysListQueryParamOrganizationIDFilterStr(str string) CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter {
	typ := CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterTypeStr

	return CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateCustomerPortalLicenseKeysListQueryParamOrganizationIDFilterArrayOfStr(arrayOfStr []string) CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter {
	typ := CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterTypeArrayOfStr

	return CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = CustomerPortalLicenseKeysListQueryParamOrganizationIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter", string(data))
}

func (u CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter: all fields are null")
}

type CustomerPortalLicenseKeysListRequest struct {
	// Filter by organization ID.
	OrganizationID *CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter `queryParam:"style=form,explode=true,name=organization_id"`
	// Filter by a specific benefit
	BenefitID *string `queryParam:"style=form,explode=true,name=benefit_id"`
	// Page number, defaults to 1.
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
}

func (c CustomerPortalLicenseKeysListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomerPortalLicenseKeysListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CustomerPortalLicenseKeysListRequest) GetOrganizationID() *CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *CustomerPortalLicenseKeysListRequest) GetBenefitID() *string {
	if o == nil {
		return nil
	}
	return o.BenefitID
}

func (o *CustomerPortalLicenseKeysListRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *CustomerPortalLicenseKeysListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type CustomerPortalLicenseKeysListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ListResourceLicenseKeyRead *components.ListResourceLicenseKeyRead

	Next func() (*CustomerPortalLicenseKeysListResponse, error)
}

func (o *CustomerPortalLicenseKeysListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CustomerPortalLicenseKeysListResponse) GetListResourceLicenseKeyRead() *components.ListResourceLicenseKeyRead {
	if o == nil {
		return nil
	}
	return o.ListResourceLicenseKeyRead
}
