// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"time"
)

type CustomFieldSelectMetadataType string

const (
	CustomFieldSelectMetadataTypeStr     CustomFieldSelectMetadataType = "str"
	CustomFieldSelectMetadataTypeInteger CustomFieldSelectMetadataType = "integer"
	CustomFieldSelectMetadataTypeBoolean CustomFieldSelectMetadataType = "boolean"
)

type CustomFieldSelectMetadata struct {
	Str     *string `queryParam:"inline"`
	Integer *int64  `queryParam:"inline"`
	Boolean *bool   `queryParam:"inline"`

	Type CustomFieldSelectMetadataType
}

func CreateCustomFieldSelectMetadataStr(str string) CustomFieldSelectMetadata {
	typ := CustomFieldSelectMetadataTypeStr

	return CustomFieldSelectMetadata{
		Str:  &str,
		Type: typ,
	}
}

func CreateCustomFieldSelectMetadataInteger(integer int64) CustomFieldSelectMetadata {
	typ := CustomFieldSelectMetadataTypeInteger

	return CustomFieldSelectMetadata{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateCustomFieldSelectMetadataBoolean(boolean bool) CustomFieldSelectMetadata {
	typ := CustomFieldSelectMetadataTypeBoolean

	return CustomFieldSelectMetadata{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *CustomFieldSelectMetadata) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CustomFieldSelectMetadataTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = CustomFieldSelectMetadataTypeInteger
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = CustomFieldSelectMetadataTypeBoolean
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CustomFieldSelectMetadata", string(data))
}

func (u CustomFieldSelectMetadata) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type CustomFieldSelectMetadata: all fields are null")
}

type CustomFieldSelectType string

const (
	CustomFieldSelectTypeSelect CustomFieldSelectType = "select"
)

func (e CustomFieldSelectType) ToPointer() *CustomFieldSelectType {
	return &e
}
func (e *CustomFieldSelectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "select":
		*e = CustomFieldSelectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CustomFieldSelectType: %v", v)
	}
}

// CustomFieldSelect - Schema for a custom field of type select.
type CustomFieldSelect struct {
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at"`
	// Last modification timestamp of the object.
	ModifiedAt *time.Time `json:"modified_at"`
	// The ID of the object.
	ID       string                               `json:"id"`
	Metadata map[string]CustomFieldSelectMetadata `json:"metadata"`
	type_    CustomFieldSelectType                `const:"select" json:"type"`
	// Identifier of the custom field. It'll be used as key when storing the value.
	Slug string `json:"slug"`
	// Name of the custom field.
	Name string `json:"name"`
	// The ID of the organization owning the custom field.
	OrganizationID string                      `json:"organization_id"`
	Properties     CustomFieldSelectProperties `json:"properties"`
}

func (c CustomFieldSelect) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomFieldSelect) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CustomFieldSelect) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *CustomFieldSelect) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *CustomFieldSelect) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CustomFieldSelect) GetMetadata() map[string]CustomFieldSelectMetadata {
	if o == nil {
		return map[string]CustomFieldSelectMetadata{}
	}
	return o.Metadata
}

func (o *CustomFieldSelect) GetType() CustomFieldSelectType {
	return CustomFieldSelectTypeSelect
}

func (o *CustomFieldSelect) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *CustomFieldSelect) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CustomFieldSelect) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *CustomFieldSelect) GetProperties() CustomFieldSelectProperties {
	if o == nil {
		return CustomFieldSelectProperties{}
	}
	return o.Properties
}
