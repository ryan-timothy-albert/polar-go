// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListResourceProduct struct {
	Items      []Product  `json:"items"`
	Pagination Pagination `json:"pagination"`
}

func (o *ListResourceProduct) GetItems() []Product {
	if o == nil {
		return []Product{}
	}
	return o.Items
}

func (o *ListResourceProduct) GetPagination() Pagination {
	if o == nil {
		return Pagination{}
	}
	return o.Pagination
}
