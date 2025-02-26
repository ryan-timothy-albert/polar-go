// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListResourceRepository struct {
	Items      []Repository `json:"items"`
	Pagination Pagination   `json:"pagination"`
}

func (o *ListResourceRepository) GetItems() []Repository {
	if o == nil {
		return []Repository{}
	}
	return o.Items
}

func (o *ListResourceRepository) GetPagination() Pagination {
	if o == nil {
		return Pagination{}
	}
	return o.Pagination
}
