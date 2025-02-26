// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"github.com/polarsource/polar-go/models/components"
)

type Oauth2RequestTokenRequestBodyType string

const (
	Oauth2RequestTokenRequestBodyTypeAuthorizationCode Oauth2RequestTokenRequestBodyType = "authorization_code"
	Oauth2RequestTokenRequestBodyTypeRefreshToken      Oauth2RequestTokenRequestBodyType = "refresh_token"
)

type Oauth2RequestTokenRequestBody struct {
	Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest *components.Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest `queryParam:"inline"`
	Onev11oauth21tokenPostXComponentsRefreshTokenRequest           *components.Onev11oauth21tokenPostXComponentsRefreshTokenRequest           `queryParam:"inline"`

	Type Oauth2RequestTokenRequestBodyType
}

func CreateOauth2RequestTokenRequestBodyAuthorizationCode(authorizationCode components.Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest) Oauth2RequestTokenRequestBody {
	typ := Oauth2RequestTokenRequestBodyTypeAuthorizationCode

	return Oauth2RequestTokenRequestBody{
		Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest: &authorizationCode,
		Type: typ,
	}
}

func CreateOauth2RequestTokenRequestBodyRefreshToken(refreshToken components.Onev11oauth21tokenPostXComponentsRefreshTokenRequest) Oauth2RequestTokenRequestBody {
	typ := Oauth2RequestTokenRequestBodyTypeRefreshToken

	return Oauth2RequestTokenRequestBody{
		Onev11oauth21tokenPostXComponentsRefreshTokenRequest: &refreshToken,
		Type: typ,
	}
}

func (u *Oauth2RequestTokenRequestBody) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		GrantType string `json:"grant_type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.GrantType {
	case "authorization_code":
		onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest := new(components.Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest)
		if err := utils.UnmarshalJSON(data, &onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (GrantType == authorization_code) type components.Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest within Oauth2RequestTokenRequestBody: %w", string(data), err)
		}

		u.Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest = onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest
		u.Type = Oauth2RequestTokenRequestBodyTypeAuthorizationCode
		return nil
	case "refresh_token":
		onev11oauth21tokenPostXComponentsRefreshTokenRequest := new(components.Onev11oauth21tokenPostXComponentsRefreshTokenRequest)
		if err := utils.UnmarshalJSON(data, &onev11oauth21tokenPostXComponentsRefreshTokenRequest, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (GrantType == refresh_token) type components.Onev11oauth21tokenPostXComponentsRefreshTokenRequest within Oauth2RequestTokenRequestBody: %w", string(data), err)
		}

		u.Onev11oauth21tokenPostXComponentsRefreshTokenRequest = onev11oauth21tokenPostXComponentsRefreshTokenRequest
		u.Type = Oauth2RequestTokenRequestBodyTypeRefreshToken
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Oauth2RequestTokenRequestBody", string(data))
}

func (u Oauth2RequestTokenRequestBody) MarshalJSON() ([]byte, error) {
	if u.Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest != nil {
		return utils.MarshalJSON(u.Onev11oauth21tokenPostXComponentsAuthorizationCodeTokenRequest, "", true)
	}

	if u.Onev11oauth21tokenPostXComponentsRefreshTokenRequest != nil {
		return utils.MarshalJSON(u.Onev11oauth21tokenPostXComponentsRefreshTokenRequest, "", true)
	}

	return nil, errors.New("could not marshal union type Oauth2RequestTokenRequestBody: all fields are null")
}

type Oauth2RequestTokenResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	TokenResponse *components.TokenResponse
}

func (o *Oauth2RequestTokenResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *Oauth2RequestTokenResponse) GetTokenResponse() *components.TokenResponse {
	if o == nil {
		return nil
	}
	return o.TokenResponse
}
