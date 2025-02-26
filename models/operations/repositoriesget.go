// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/polarsource/polar-go/models/components"
)

type RepositoriesGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RepositoriesGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RepositoriesGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Repository *components.Repository
}

func (o *RepositoriesGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RepositoriesGetResponse) GetRepository() *components.Repository {
	if o == nil {
		return nil
	}
	return o.Repository
}
