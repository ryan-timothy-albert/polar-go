// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListResourceCustomerBenefitGrant struct {
	Items      []CustomerBenefitGrant `json:"items"`
	Pagination Pagination             `json:"pagination"`
}

func (o *ListResourceCustomerBenefitGrant) GetItems() []CustomerBenefitGrant {
	if o == nil {
		return []CustomerBenefitGrant{}
	}
	return o.Items
}

func (o *ListResourceCustomerBenefitGrant) GetPagination() Pagination {
	if o == nil {
		return Pagination{}
	}
	return o.Pagination
}
