// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/polarsource/polar-go/models/components"
)

type CheckoutsCustomGetRequest struct {
	// The checkout session ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *CheckoutsCustomGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type CheckoutsCustomGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Checkout *components.Checkout
}

func (o *CheckoutsCustomGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CheckoutsCustomGetResponse) GetCheckout() *components.Checkout {
	if o == nil {
		return nil
	}
	return o.Checkout
}
