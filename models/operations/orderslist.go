// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"github.com/polarsource/polar-go/models/components"
)

type OrdersListQueryParamOrganizationIDFilterType string

const (
	OrdersListQueryParamOrganizationIDFilterTypeStr        OrdersListQueryParamOrganizationIDFilterType = "str"
	OrdersListQueryParamOrganizationIDFilterTypeArrayOfStr OrdersListQueryParamOrganizationIDFilterType = "arrayOfStr"
)

// OrdersListQueryParamOrganizationIDFilter - Filter by organization ID.
type OrdersListQueryParamOrganizationIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type OrdersListQueryParamOrganizationIDFilterType
}

func CreateOrdersListQueryParamOrganizationIDFilterStr(str string) OrdersListQueryParamOrganizationIDFilter {
	typ := OrdersListQueryParamOrganizationIDFilterTypeStr

	return OrdersListQueryParamOrganizationIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateOrdersListQueryParamOrganizationIDFilterArrayOfStr(arrayOfStr []string) OrdersListQueryParamOrganizationIDFilter {
	typ := OrdersListQueryParamOrganizationIDFilterTypeArrayOfStr

	return OrdersListQueryParamOrganizationIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *OrdersListQueryParamOrganizationIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = OrdersListQueryParamOrganizationIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = OrdersListQueryParamOrganizationIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for OrdersListQueryParamOrganizationIDFilter", string(data))
}

func (u OrdersListQueryParamOrganizationIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type OrdersListQueryParamOrganizationIDFilter: all fields are null")
}

type QueryParamProductIDFilterType string

const (
	QueryParamProductIDFilterTypeStr        QueryParamProductIDFilterType = "str"
	QueryParamProductIDFilterTypeArrayOfStr QueryParamProductIDFilterType = "arrayOfStr"
)

// QueryParamProductIDFilter - Filter by product ID.
type QueryParamProductIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type QueryParamProductIDFilterType
}

func CreateQueryParamProductIDFilterStr(str string) QueryParamProductIDFilter {
	typ := QueryParamProductIDFilterTypeStr

	return QueryParamProductIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateQueryParamProductIDFilterArrayOfStr(arrayOfStr []string) QueryParamProductIDFilter {
	typ := QueryParamProductIDFilterTypeArrayOfStr

	return QueryParamProductIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *QueryParamProductIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = QueryParamProductIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = QueryParamProductIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for QueryParamProductIDFilter", string(data))
}

func (u QueryParamProductIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type QueryParamProductIDFilter: all fields are null")
}

type ProductPriceTypeFilterType string

const (
	ProductPriceTypeFilterTypeProductPriceType        ProductPriceTypeFilterType = "ProductPriceType"
	ProductPriceTypeFilterTypeArrayOfProductPriceType ProductPriceTypeFilterType = "arrayOfProductPriceType"
)

// ProductPriceTypeFilter - Filter by product price type. `recurring` will return orders corresponding to subscriptions creations or renewals. `one_time` will return orders corresponding to one-time purchases.
type ProductPriceTypeFilter struct {
	ProductPriceType        *components.ProductPriceType  `queryParam:"inline"`
	ArrayOfProductPriceType []components.ProductPriceType `queryParam:"inline"`

	Type ProductPriceTypeFilterType
}

func CreateProductPriceTypeFilterProductPriceType(productPriceType components.ProductPriceType) ProductPriceTypeFilter {
	typ := ProductPriceTypeFilterTypeProductPriceType

	return ProductPriceTypeFilter{
		ProductPriceType: &productPriceType,
		Type:             typ,
	}
}

func CreateProductPriceTypeFilterArrayOfProductPriceType(arrayOfProductPriceType []components.ProductPriceType) ProductPriceTypeFilter {
	typ := ProductPriceTypeFilterTypeArrayOfProductPriceType

	return ProductPriceTypeFilter{
		ArrayOfProductPriceType: arrayOfProductPriceType,
		Type:                    typ,
	}
}

func (u *ProductPriceTypeFilter) UnmarshalJSON(data []byte) error {

	var productPriceType components.ProductPriceType = components.ProductPriceType("")
	if err := utils.UnmarshalJSON(data, &productPriceType, "", true, true); err == nil {
		u.ProductPriceType = &productPriceType
		u.Type = ProductPriceTypeFilterTypeProductPriceType
		return nil
	}

	var arrayOfProductPriceType []components.ProductPriceType = []components.ProductPriceType{}
	if err := utils.UnmarshalJSON(data, &arrayOfProductPriceType, "", true, true); err == nil {
		u.ArrayOfProductPriceType = arrayOfProductPriceType
		u.Type = ProductPriceTypeFilterTypeArrayOfProductPriceType
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ProductPriceTypeFilter", string(data))
}

func (u ProductPriceTypeFilter) MarshalJSON() ([]byte, error) {
	if u.ProductPriceType != nil {
		return utils.MarshalJSON(u.ProductPriceType, "", true)
	}

	if u.ArrayOfProductPriceType != nil {
		return utils.MarshalJSON(u.ArrayOfProductPriceType, "", true)
	}

	return nil, errors.New("could not marshal union type ProductPriceTypeFilter: all fields are null")
}

type QueryParamDiscountIDFilterType string

const (
	QueryParamDiscountIDFilterTypeStr        QueryParamDiscountIDFilterType = "str"
	QueryParamDiscountIDFilterTypeArrayOfStr QueryParamDiscountIDFilterType = "arrayOfStr"
)

// QueryParamDiscountIDFilter - Filter by discount ID.
type QueryParamDiscountIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type QueryParamDiscountIDFilterType
}

func CreateQueryParamDiscountIDFilterStr(str string) QueryParamDiscountIDFilter {
	typ := QueryParamDiscountIDFilterTypeStr

	return QueryParamDiscountIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateQueryParamDiscountIDFilterArrayOfStr(arrayOfStr []string) QueryParamDiscountIDFilter {
	typ := QueryParamDiscountIDFilterTypeArrayOfStr

	return QueryParamDiscountIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *QueryParamDiscountIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = QueryParamDiscountIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = QueryParamDiscountIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for QueryParamDiscountIDFilter", string(data))
}

func (u QueryParamDiscountIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type QueryParamDiscountIDFilter: all fields are null")
}

type OrdersListQueryParamCustomerIDFilterType string

const (
	OrdersListQueryParamCustomerIDFilterTypeStr        OrdersListQueryParamCustomerIDFilterType = "str"
	OrdersListQueryParamCustomerIDFilterTypeArrayOfStr OrdersListQueryParamCustomerIDFilterType = "arrayOfStr"
)

// OrdersListQueryParamCustomerIDFilter - Filter by customer ID.
type OrdersListQueryParamCustomerIDFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type OrdersListQueryParamCustomerIDFilterType
}

func CreateOrdersListQueryParamCustomerIDFilterStr(str string) OrdersListQueryParamCustomerIDFilter {
	typ := OrdersListQueryParamCustomerIDFilterTypeStr

	return OrdersListQueryParamCustomerIDFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateOrdersListQueryParamCustomerIDFilterArrayOfStr(arrayOfStr []string) OrdersListQueryParamCustomerIDFilter {
	typ := OrdersListQueryParamCustomerIDFilterTypeArrayOfStr

	return OrdersListQueryParamCustomerIDFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *OrdersListQueryParamCustomerIDFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = OrdersListQueryParamCustomerIDFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = OrdersListQueryParamCustomerIDFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for OrdersListQueryParamCustomerIDFilter", string(data))
}

func (u OrdersListQueryParamCustomerIDFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type OrdersListQueryParamCustomerIDFilter: all fields are null")
}

type OrdersListRequest struct {
	// Filter by organization ID.
	OrganizationID *OrdersListQueryParamOrganizationIDFilter `queryParam:"style=form,explode=true,name=organization_id"`
	// Filter by product ID.
	ProductID *QueryParamProductIDFilter `queryParam:"style=form,explode=true,name=product_id"`
	// Filter by product price type. `recurring` will return orders corresponding to subscriptions creations or renewals. `one_time` will return orders corresponding to one-time purchases.
	ProductPriceType *ProductPriceTypeFilter `queryParam:"style=form,explode=true,name=product_price_type"`
	// Filter by discount ID.
	DiscountID *QueryParamDiscountIDFilter `queryParam:"style=form,explode=true,name=discount_id"`
	// Filter by customer ID.
	CustomerID *OrdersListQueryParamCustomerIDFilter `queryParam:"style=form,explode=true,name=customer_id"`
	// Page number, defaults to 1.
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Size of a page, defaults to 10. Maximum is 100.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// Sorting criterion. Several criteria can be used simultaneously and will be applied in order. Add a minus sign `-` before the criteria name to sort by descending order.
	Sorting []components.OrderSortProperty `queryParam:"style=form,explode=true,name=sorting"`
}

func (o OrdersListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OrdersListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OrdersListRequest) GetOrganizationID() *OrdersListQueryParamOrganizationIDFilter {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *OrdersListRequest) GetProductID() *QueryParamProductIDFilter {
	if o == nil {
		return nil
	}
	return o.ProductID
}

func (o *OrdersListRequest) GetProductPriceType() *ProductPriceTypeFilter {
	if o == nil {
		return nil
	}
	return o.ProductPriceType
}

func (o *OrdersListRequest) GetDiscountID() *QueryParamDiscountIDFilter {
	if o == nil {
		return nil
	}
	return o.DiscountID
}

func (o *OrdersListRequest) GetCustomerID() *OrdersListQueryParamCustomerIDFilter {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *OrdersListRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *OrdersListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *OrdersListRequest) GetSorting() []components.OrderSortProperty {
	if o == nil {
		return nil
	}
	return o.Sorting
}

type OrdersListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ListResourceOrder *components.ListResourceOrder

	Next func() (*OrdersListResponse, error)
}

func (o *OrdersListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OrdersListResponse) GetListResourceOrder() *components.ListResourceOrder {
	if o == nil {
		return nil
	}
	return o.ListResourceOrder
}
