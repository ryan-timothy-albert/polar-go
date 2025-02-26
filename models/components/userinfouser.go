// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UserInfoUser struct {
	Sub           string  `json:"sub"`
	Name          *string `json:"name,omitempty"`
	Email         *string `json:"email,omitempty"`
	EmailVerified *bool   `json:"email_verified,omitempty"`
}

func (o *UserInfoUser) GetSub() string {
	if o == nil {
		return ""
	}
	return o.Sub
}

func (o *UserInfoUser) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UserInfoUser) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *UserInfoUser) GetEmailVerified() *bool {
	if o == nil {
		return nil
	}
	return o.EmailVerified
}
