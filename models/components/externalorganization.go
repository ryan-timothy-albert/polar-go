// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/polarsource/polar-go/internal/utils"
)

type ExternalOrganization struct {
	ID              string    `json:"id"`
	platform        Platforms `const:"github" json:"platform"`
	Name            string    `json:"name"`
	AvatarURL       string    `json:"avatar_url"`
	IsPersonal      bool      `json:"is_personal"`
	Bio             *string   `json:"bio"`
	PrettyName      *string   `json:"pretty_name"`
	Company         *string   `json:"company"`
	Blog            *string   `json:"blog"`
	Location        *string   `json:"location"`
	Email           *string   `json:"email"`
	TwitterUsername *string   `json:"twitter_username"`
	OrganizationID  *string   `json:"organization_id"`
}

func (e ExternalOrganization) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExternalOrganization) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ExternalOrganization) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ExternalOrganization) GetPlatform() Platforms {
	return PlatformsGithub
}

func (o *ExternalOrganization) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ExternalOrganization) GetAvatarURL() string {
	if o == nil {
		return ""
	}
	return o.AvatarURL
}

func (o *ExternalOrganization) GetIsPersonal() bool {
	if o == nil {
		return false
	}
	return o.IsPersonal
}

func (o *ExternalOrganization) GetBio() *string {
	if o == nil {
		return nil
	}
	return o.Bio
}

func (o *ExternalOrganization) GetPrettyName() *string {
	if o == nil {
		return nil
	}
	return o.PrettyName
}

func (o *ExternalOrganization) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *ExternalOrganization) GetBlog() *string {
	if o == nil {
		return nil
	}
	return o.Blog
}

func (o *ExternalOrganization) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *ExternalOrganization) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ExternalOrganization) GetTwitterUsername() *string {
	if o == nil {
		return nil
	}
	return o.TwitterUsername
}

func (o *ExternalOrganization) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}
