// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package webhooks

import (
	"github.com/polarsource/polar-go/models/components"
)

type EndpointbenefitGrantRevokedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (o *EndpointbenefitGrantRevokedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *EndpointbenefitGrantRevokedPostResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}
