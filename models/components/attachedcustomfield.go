// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AttachedCustomField - Schema of a custom field attached to a resource.
type AttachedCustomField struct {
	// ID of the custom field.
	CustomFieldID string      `json:"custom_field_id"`
	CustomField   CustomField `json:"custom_field"`
	// Order of the custom field in the resource.
	Order int64 `json:"order"`
	// Whether the value is required for this custom field.
	Required bool `json:"required"`
}

func (o *AttachedCustomField) GetCustomFieldID() string {
	if o == nil {
		return ""
	}
	return o.CustomFieldID
}

func (o *AttachedCustomField) GetCustomField() CustomField {
	if o == nil {
		return CustomField{}
	}
	return o.CustomField
}

func (o *AttachedCustomField) GetCustomFieldCheckbox() *CustomFieldCheckbox {
	return o.GetCustomField().CustomFieldCheckbox
}

func (o *AttachedCustomField) GetCustomFieldDate() *CustomFieldDate {
	return o.GetCustomField().CustomFieldDate
}

func (o *AttachedCustomField) GetCustomFieldNumber() *CustomFieldNumber {
	return o.GetCustomField().CustomFieldNumber
}

func (o *AttachedCustomField) GetCustomFieldSelect() *CustomFieldSelect {
	return o.GetCustomField().CustomFieldSelect
}

func (o *AttachedCustomField) GetCustomFieldText() *CustomFieldText {
	return o.GetCustomField().CustomFieldText
}

func (o *AttachedCustomField) GetOrder() int64 {
	if o == nil {
		return 0
	}
	return o.Order
}

func (o *AttachedCustomField) GetRequired() bool {
	if o == nil {
		return false
	}
	return o.Required
}
