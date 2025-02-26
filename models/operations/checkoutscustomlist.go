// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"github.com/polarsource/polar-go/models/components"
)

type CheckoutsCustomListQueryParamOrganizationIDFilterType string

const (
	CheckoutsCustomListQueryParamOrganizationIDFilterTypeStr        CheckoutsCustomListQueryParamOrganizationIDFilterType = "str"
	CheckoutsCustomListQueryParamOrganizationIDFilterTypeArrayOfStr CheckoutsCustomListQueryParamOrganizationIDFilterType = "arrayOfStr"
)

// CheckoutsCustomListQueryParamOrganizationIDFilter - Filter by organization ID.
type CheckoutsCustomListQueryParamOrganizationIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type CheckoutsCustomListQueryParamOrganizationIDFilterType
}

func CreateCheckoutsCustomListQueryParamOrganizationIDFilterStr(str string) CheckoutsCustomListQueryParamOrganizationIDFilter {
	typ := CheckoutsCustomListQueryParamOrganizationIDFilterTypeStr

	return CheckoutsCustomListQueryParamOrganizationIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateCheckoutsCustomListQueryParamOrganizationIDFilterArrayOfStr(arrayOfStr []string) CheckoutsCustomListQueryParamOrganizationIDFilter {
	typ := CheckoutsCustomListQueryParamOrganizationIDFilterTypeArrayOfStr

	return CheckoutsCustomListQueryParamOrganizationIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *CheckoutsCustomListQueryParamOrganizationIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CheckoutsCustomListQueryParamOrganizationIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = CheckoutsCustomListQueryParamOrganizationIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CheckoutsCustomListQueryParamOrganizationIDFilter", string(data))
}

func (u CheckoutsCustomListQueryParamOrganizationIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type CheckoutsCustomListQueryParamOrganizationIDFilter: all fields are null")
}

type CheckoutsCustomListQueryParamProductIDFilterType string

const (
	CheckoutsCustomListQueryParamProductIDFilterTypeStr        CheckoutsCustomListQueryParamProductIDFilterType = "str"
	CheckoutsCustomListQueryParamProductIDFilterTypeArrayOfStr CheckoutsCustomListQueryParamProductIDFilterType = "arrayOfStr"
)

// CheckoutsCustomListQueryParamProductIDFilter - Filter by product ID.
type CheckoutsCustomListQueryParamProductIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type CheckoutsCustomListQueryParamProductIDFilterType
}

func CreateCheckoutsCustomListQueryParamProductIDFilterStr(str string) CheckoutsCustomListQueryParamProductIDFilter {
	typ := CheckoutsCustomListQueryParamProductIDFilterTypeStr

	return CheckoutsCustomListQueryParamProductIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateCheckoutsCustomListQueryParamProductIDFilterArrayOfStr(arrayOfStr []string) CheckoutsCustomListQueryParamProductIDFilter {
	typ := CheckoutsCustomListQueryParamProductIDFilterTypeArrayOfStr

	return CheckoutsCustomListQueryParamProductIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *CheckoutsCustomListQueryParamProductIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CheckoutsCustomListQueryParamProductIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = CheckoutsCustomListQueryParamProductIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CheckoutsCustomListQueryParamProductIDFilter", string(data))
}

func (u CheckoutsCustomListQueryParamProductIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type CheckoutsCustomListQueryParamProductIDFilter: all fields are null")
}

type CheckoutsCustomListRequest struct {
	// Filter by organization ID.
	OrganizationID *CheckoutsCustomListQueryParamOrganizationIDFilter `queryParam:"style=form,explode=true,name=organization_id"`
	// Filter by product ID.
	ProductID *CheckoutsCustomListQueryParamProductIDFilter `queryParam:"style=form,explode=true,name=product_id"`
	// Page number, defaults to 1.
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// Sorting criterion. Several criteria can be used simultaneously and will be applied in order. Add a minus sign `-` before the criteria name to sort by descending order.
	Sorting []components.CheckoutSortProperty `queryParam:"style=form,explode=true,name=sorting"`
}

func (c CheckoutsCustomListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CheckoutsCustomListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CheckoutsCustomListRequest) GetOrganizationID() *CheckoutsCustomListQueryParamOrganizationIDFilter {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *CheckoutsCustomListRequest) GetProductID() *CheckoutsCustomListQueryParamProductIDFilter {
	if o == nil {
		return nil
	}
	return o.ProductID
}

func (o *CheckoutsCustomListRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *CheckoutsCustomListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *CheckoutsCustomListRequest) GetSorting() []components.CheckoutSortProperty {
	if o == nil {
		return nil
	}
	return o.Sorting
}

type CheckoutsCustomListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ListResourceCheckout *components.ListResourceCheckout

	Next func() (*CheckoutsCustomListResponse, error)
}

func (o *CheckoutsCustomListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CheckoutsCustomListResponse) GetListResourceCheckout() *components.ListResourceCheckout {
	if o == nil {
		return nil
	}
	return o.ListResourceCheckout
}
