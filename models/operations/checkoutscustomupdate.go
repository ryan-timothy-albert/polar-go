// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/polarsource/polar-go/models/components"
)

type CheckoutsCustomUpdateRequest struct {
	// The checkout session ID.
	ID             string                    `pathParam:"style=simple,explode=false,name=id"`
	CheckoutUpdate components.CheckoutUpdate `request:"mediaType=application/json"`
}

func (o *CheckoutsCustomUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CheckoutsCustomUpdateRequest) GetCheckoutUpdate() components.CheckoutUpdate {
	if o == nil {
		return components.CheckoutUpdate{}
	}
	return o.CheckoutUpdate
}

type CheckoutsCustomUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Checkout session updated.
	Checkout *components.Checkout
}

func (o *CheckoutsCustomUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CheckoutsCustomUpdateResponse) GetCheckout() *components.Checkout {
	if o == nil {
		return nil
	}
	return o.Checkout
}
