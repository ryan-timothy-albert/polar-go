// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/polarsource/polar-go/internal/utils"
	"time"
)

type LicenseKeyWithActivations struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	UserID     string             `json:"user_id"`
	CustomerID string             `json:"customer_id"`
	User       LicenseKeyUser     `json:"user"`
	Customer   LicenseKeyCustomer `json:"customer"`
	// The benefit ID.
	BenefitID        string                     `json:"benefit_id"`
	Key              string                     `json:"key"`
	DisplayKey       string                     `json:"display_key"`
	Status           LicenseKeyStatus           `json:"status"`
	LimitActivations *int64                     `json:"limit_activations"`
	Usage            int64                      `json:"usage"`
	LimitUsage       *int64                     `json:"limit_usage"`
	Validations      int64                      `json:"validations"`
	LastValidatedAt  *time.Time                 `json:"last_validated_at"`
	ExpiresAt        *time.Time                 `json:"expires_at"`
	Activations      []LicenseKeyActivationBase `json:"activations"`
}

func (l LicenseKeyWithActivations) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LicenseKeyWithActivations) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LicenseKeyWithActivations) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LicenseKeyWithActivations) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *LicenseKeyWithActivations) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *LicenseKeyWithActivations) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *LicenseKeyWithActivations) GetUser() LicenseKeyUser {
	if o == nil {
		return LicenseKeyUser{}
	}
	return o.User
}

func (o *LicenseKeyWithActivations) GetCustomer() LicenseKeyCustomer {
	if o == nil {
		return LicenseKeyCustomer{}
	}
	return o.Customer
}

func (o *LicenseKeyWithActivations) GetBenefitID() string {
	if o == nil {
		return ""
	}
	return o.BenefitID
}

func (o *LicenseKeyWithActivations) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *LicenseKeyWithActivations) GetDisplayKey() string {
	if o == nil {
		return ""
	}
	return o.DisplayKey
}

func (o *LicenseKeyWithActivations) GetStatus() LicenseKeyStatus {
	if o == nil {
		return LicenseKeyStatus("")
	}
	return o.Status
}

func (o *LicenseKeyWithActivations) GetLimitActivations() *int64 {
	if o == nil {
		return nil
	}
	return o.LimitActivations
}

func (o *LicenseKeyWithActivations) GetUsage() int64 {
	if o == nil {
		return 0
	}
	return o.Usage
}

func (o *LicenseKeyWithActivations) GetLimitUsage() *int64 {
	if o == nil {
		return nil
	}
	return o.LimitUsage
}

func (o *LicenseKeyWithActivations) GetValidations() int64 {
	if o == nil {
		return 0
	}
	return o.Validations
}

func (o *LicenseKeyWithActivations) GetLastValidatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastValidatedAt
}

func (o *LicenseKeyWithActivations) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *LicenseKeyWithActivations) GetActivations() []LicenseKeyActivationBase {
	if o == nil {
		return []LicenseKeyActivationBase{}
	}
	return o.Activations
}
