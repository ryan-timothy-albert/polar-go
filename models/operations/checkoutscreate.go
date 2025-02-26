// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/polarsource/polar-go/models/components"
)

type CheckoutsCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	CheckoutLegacy *components.CheckoutLegacy
}

func (o *CheckoutsCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CheckoutsCreateResponse) GetCheckoutLegacy() *components.CheckoutLegacy {
	if o == nil {
		return nil
	}
	return o.CheckoutLegacy
}
