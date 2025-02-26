// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Conditions struct {
}

type LicenseKeyValidate struct {
	Key            string      `json:"key"`
	OrganizationID string      `json:"organization_id"`
	ActivationID   *string     `json:"activation_id,omitempty"`
	BenefitID      *string     `json:"benefit_id,omitempty"`
	CustomerID     *string     `json:"customer_id,omitempty"`
	IncrementUsage *int64      `json:"increment_usage,omitempty"`
	Conditions     *Conditions `json:"conditions,omitempty"`
}

func (o *LicenseKeyValidate) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *LicenseKeyValidate) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *LicenseKeyValidate) GetActivationID() *string {
	if o == nil {
		return nil
	}
	return o.ActivationID
}

func (o *LicenseKeyValidate) GetBenefitID() *string {
	if o == nil {
		return nil
	}
	return o.BenefitID
}

func (o *LicenseKeyValidate) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *LicenseKeyValidate) GetIncrementUsage() *int64 {
	if o == nil {
		return nil
	}
	return o.IncrementUsage
}

func (o *LicenseKeyValidate) GetConditions() *Conditions {
	if o == nil {
		return nil
	}
	return o.Conditions
}
