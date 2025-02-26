// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
)

type CustomFieldUpdateTextMetadataType string

const (
	CustomFieldUpdateTextMetadataTypeStr     CustomFieldUpdateTextMetadataType = "str"
	CustomFieldUpdateTextMetadataTypeInteger CustomFieldUpdateTextMetadataType = "integer"
	CustomFieldUpdateTextMetadataTypeBoolean CustomFieldUpdateTextMetadataType = "boolean"
)

type CustomFieldUpdateTextMetadata struct {
	Str     *string `queryParam:"inline"`
	Integer *int64  `queryParam:"inline"`
	Boolean *bool   `queryParam:"inline"`

	Type CustomFieldUpdateTextMetadataType
}

func CreateCustomFieldUpdateTextMetadataStr(str string) CustomFieldUpdateTextMetadata {
	typ := CustomFieldUpdateTextMetadataTypeStr

	return CustomFieldUpdateTextMetadata{
		Str:  &str,
		Type: typ,
	}
}

func CreateCustomFieldUpdateTextMetadataInteger(integer int64) CustomFieldUpdateTextMetadata {
	typ := CustomFieldUpdateTextMetadataTypeInteger

	return CustomFieldUpdateTextMetadata{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateCustomFieldUpdateTextMetadataBoolean(boolean bool) CustomFieldUpdateTextMetadata {
	typ := CustomFieldUpdateTextMetadataTypeBoolean

	return CustomFieldUpdateTextMetadata{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *CustomFieldUpdateTextMetadata) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CustomFieldUpdateTextMetadataTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = CustomFieldUpdateTextMetadataTypeInteger
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = CustomFieldUpdateTextMetadataTypeBoolean
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CustomFieldUpdateTextMetadata", string(data))
}

func (u CustomFieldUpdateTextMetadata) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type CustomFieldUpdateTextMetadata: all fields are null")
}

type CustomFieldUpdateTextType string

const (
	CustomFieldUpdateTextTypeText CustomFieldUpdateTextType = "text"
)

func (e CustomFieldUpdateTextType) ToPointer() *CustomFieldUpdateTextType {
	return &e
}
func (e *CustomFieldUpdateTextType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = CustomFieldUpdateTextType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CustomFieldUpdateTextType: %v", v)
	}
}

// CustomFieldUpdateText - Schema to update a custom field of type text.
type CustomFieldUpdateText struct {
	Metadata   map[string]CustomFieldUpdateTextMetadata `json:"metadata,omitempty"`
	Name       *string                                  `json:"name,omitempty"`
	Slug       *string                                  `json:"slug,omitempty"`
	type_      CustomFieldUpdateTextType                `const:"text" json:"type"`
	Properties *CustomFieldTextProperties               `json:"properties,omitempty"`
}

func (c CustomFieldUpdateText) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomFieldUpdateText) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CustomFieldUpdateText) GetMetadata() map[string]CustomFieldUpdateTextMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CustomFieldUpdateText) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CustomFieldUpdateText) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CustomFieldUpdateText) GetType() CustomFieldUpdateTextType {
	return CustomFieldUpdateTextTypeText
}

func (o *CustomFieldUpdateText) GetProperties() *CustomFieldTextProperties {
	if o == nil {
		return nil
	}
	return o.Properties
}
