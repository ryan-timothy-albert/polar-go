// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/polarsource/polar-go/internal/utils"
	"time"
)

type LicenseKeyActivationReadMeta struct {
}

type LicenseKeyActivationRead struct {
	ID           string                       `json:"id"`
	LicenseKeyID string                       `json:"license_key_id"`
	Label        string                       `json:"label"`
	Meta         LicenseKeyActivationReadMeta `json:"meta"`
	CreatedAt    time.Time                    `json:"created_at"`
	ModifiedAt   *time.Time                   `json:"modified_at"`
	LicenseKey   LicenseKeyRead               `json:"license_key"`
}

func (l LicenseKeyActivationRead) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LicenseKeyActivationRead) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LicenseKeyActivationRead) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LicenseKeyActivationRead) GetLicenseKeyID() string {
	if o == nil {
		return ""
	}
	return o.LicenseKeyID
}

func (o *LicenseKeyActivationRead) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *LicenseKeyActivationRead) GetMeta() LicenseKeyActivationReadMeta {
	if o == nil {
		return LicenseKeyActivationReadMeta{}
	}
	return o.Meta
}

func (o *LicenseKeyActivationRead) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *LicenseKeyActivationRead) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *LicenseKeyActivationRead) GetLicenseKey() LicenseKeyRead {
	if o == nil {
		return LicenseKeyRead{}
	}
	return o.LicenseKey
}
