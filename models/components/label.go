// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (o *Label) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Label) GetColor() string {
	if o == nil {
		return ""
	}
	return o.Color
}
