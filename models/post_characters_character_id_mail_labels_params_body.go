package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCharactersCharacterIDMailLabelsParamsBody post_characters_character_id_mail_labels_label
//
// label object
// swagger:model postCharactersCharacterIdMailLabelsParamsBody
type PostCharactersCharacterIDMailLabelsParamsBody struct {

	// post_characters_character_id_mail_labels_color
	//
	// Hexadecimal string representing label color,
	// in RGB format
	//
	Color *string `json:"color,omitempty"`

	// post_characters_character_id_mail_labels_name
	//
	// name string
	// Required: true
	// Max Length: 40
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this post characters character Id mail labels params body
func (m *PostCharactersCharacterIDMailLabelsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postCharactersCharacterIdMailLabelsParamsBodyTypeColorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["#ffffff","#ffff01","#ff6600","#fe0000","#9a0000","#660066","#0000fe","#0099ff","#01ffff","#00ff33","#349800","#006634","#666666","#999999","#e6e6e6","#ffffcd","#99ffff","#ccff9a"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCharactersCharacterIdMailLabelsParamsBodyTypeColorPropEnum = append(postCharactersCharacterIdMailLabelsParamsBodyTypeColorPropEnum, v)
	}
}

const (
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNrFfffff captures enum value "#ffffff"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNrFfffff string = "#ffffff"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNrFfff01 captures enum value "#ffff01"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNrFfff01 string = "#ffff01"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNrFf6600 captures enum value "#ff6600"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNrFf6600 string = "#ff6600"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNrFe0000 captures enum value "#fe0000"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNrFe0000 string = "#fe0000"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr9a0000 captures enum value "#9a0000"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr9a0000 string = "#9a0000"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr660066 captures enum value "#660066"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr660066 string = "#660066"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr0000fe captures enum value "#0000fe"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr0000fe string = "#0000fe"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr0099ff captures enum value "#0099ff"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr0099ff string = "#0099ff"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr01ffff captures enum value "#01ffff"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr01ffff string = "#01ffff"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr00ff33 captures enum value "#00ff33"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr00ff33 string = "#00ff33"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr349800 captures enum value "#349800"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr349800 string = "#349800"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr006634 captures enum value "#006634"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr006634 string = "#006634"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr666666 captures enum value "#666666"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr666666 string = "#666666"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr999999 captures enum value "#999999"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr999999 string = "#999999"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNrE6e6e6 captures enum value "#e6e6e6"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNrE6e6e6 string = "#e6e6e6"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNrFfffcd captures enum value "#ffffcd"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNrFfffcd string = "#ffffcd"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNr99ffff captures enum value "#99ffff"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNr99ffff string = "#99ffff"
	// PostCharactersCharacterIDMailLabelsParamsBodyColorNrCcff9a captures enum value "#ccff9a"
	PostCharactersCharacterIDMailLabelsParamsBodyColorNrCcff9a string = "#ccff9a"
)

// prop value enum
func (m *PostCharactersCharacterIDMailLabelsParamsBody) validateColorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postCharactersCharacterIdMailLabelsParamsBodyTypeColorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostCharactersCharacterIDMailLabelsParamsBody) validateColor(formats strfmt.Registry) error {

	if swag.IsZero(m.Color) { // not required
		return nil
	}

	// value enum
	if err := m.validateColorEnum("color", "body", *m.Color); err != nil {
		return err
	}

	return nil
}

func (m *PostCharactersCharacterIDMailLabelsParamsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersCharacterIDMailLabelsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersCharacterIDMailLabelsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDMailLabelsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
