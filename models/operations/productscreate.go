// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/polarsource/polar-go/models/components"
)

type ProductsCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Product created.
	Product *components.Product
}

func (o *ProductsCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ProductsCreateResponse) GetProduct() *components.Product {
	if o == nil {
		return nil
	}
	return o.Product
}
