// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"github.com/polarsource/polar-go/models/components"
)

type ProductsListQueryParamOrganizationIDFilterType string

const (
	ProductsListQueryParamOrganizationIDFilterTypeStr        ProductsListQueryParamOrganizationIDFilterType = "str"
	ProductsListQueryParamOrganizationIDFilterTypeArrayOfStr ProductsListQueryParamOrganizationIDFilterType = "arrayOfStr"
)

// ProductsListQueryParamOrganizationIDFilter - Filter by organization ID.
type ProductsListQueryParamOrganizationIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type ProductsListQueryParamOrganizationIDFilterType
}

func CreateProductsListQueryParamOrganizationIDFilterStr(str string) ProductsListQueryParamOrganizationIDFilter {
	typ := ProductsListQueryParamOrganizationIDFilterTypeStr

	return ProductsListQueryParamOrganizationIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateProductsListQueryParamOrganizationIDFilterArrayOfStr(arrayOfStr []string) ProductsListQueryParamOrganizationIDFilter {
	typ := ProductsListQueryParamOrganizationIDFilterTypeArrayOfStr

	return ProductsListQueryParamOrganizationIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *ProductsListQueryParamOrganizationIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ProductsListQueryParamOrganizationIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = ProductsListQueryParamOrganizationIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ProductsListQueryParamOrganizationIDFilter", string(data))
}

func (u ProductsListQueryParamOrganizationIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type ProductsListQueryParamOrganizationIDFilter: all fields are null")
}

type BenefitIDFilterType string

const (
	BenefitIDFilterTypeStr        BenefitIDFilterType = "str"
	BenefitIDFilterTypeArrayOfStr BenefitIDFilterType = "arrayOfStr"
)

// BenefitIDFilter - Filter products granting specific benefit.
type BenefitIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type BenefitIDFilterType
}

func CreateBenefitIDFilterStr(str string) BenefitIDFilter {
	typ := BenefitIDFilterTypeStr

	return BenefitIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateBenefitIDFilterArrayOfStr(arrayOfStr []string) BenefitIDFilter {
	typ := BenefitIDFilterTypeArrayOfStr

	return BenefitIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *BenefitIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = BenefitIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = BenefitIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for BenefitIDFilter", string(data))
}

func (u BenefitIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type BenefitIDFilter: all fields are null")
}

type ProductsListRequest struct {
	// Filter by organization ID.
	OrganizationID *ProductsListQueryParamOrganizationIDFilter `queryParam:"style=form,explode=true,name=organization_id"`
	// Filter by product name.
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Filter on archived products.
	IsArchived *bool `queryParam:"style=form,explode=true,name=is_archived"`
	// Filter on recurring products. If `true`, only subscriptions tiers are returned. If `false`, only one-time purchase products are returned.
	IsRecurring *bool `queryParam:"style=form,explode=true,name=is_recurring"`
	// Filter products granting specific benefit.
	BenefitID *BenefitIDFilter `queryParam:"style=form,explode=true,name=benefit_id"`
	// Page number, defaults to 1.
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// Sorting criterion. Several criteria can be used simultaneously and will be applied in order. Add a minus sign `-` before the criteria name to sort by descending order.
	Sorting []components.ProductSortProperty `queryParam:"style=form,explode=true,name=sorting"`
}

func (p ProductsListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductsListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProductsListRequest) GetOrganizationID() *ProductsListQueryParamOrganizationIDFilter {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *ProductsListRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ProductsListRequest) GetIsArchived() *bool {
	if o == nil {
		return nil
	}
	return o.IsArchived
}

func (o *ProductsListRequest) GetIsRecurring() *bool {
	if o == nil {
		return nil
	}
	return o.IsRecurring
}

func (o *ProductsListRequest) GetBenefitID() *BenefitIDFilter {
	if o == nil {
		return nil
	}
	return o.BenefitID
}

func (o *ProductsListRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ProductsListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ProductsListRequest) GetSorting() []components.ProductSortProperty {
	if o == nil {
		return nil
	}
	return o.Sorting
}

type ProductsListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ListResourceProduct *components.ListResourceProduct

	Next func() (*ProductsListResponse, error)
}

func (o *ProductsListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ProductsListResponse) GetListResourceProduct() *components.ListResourceProduct {
	if o == nil {
		return nil
	}
	return o.ListResourceProduct
}
