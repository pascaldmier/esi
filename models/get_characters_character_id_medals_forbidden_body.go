package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDMedalsForbiddenBody get_characters_character_id_medals_forbidden
//
// Forbidden
// swagger:model getCharactersCharacterIdMedalsForbiddenBody
type GetCharactersCharacterIDMedalsForbiddenBody struct {

	// get_characters_character_id_medals_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id medals forbidden body
func (m *GetCharactersCharacterIDMedalsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDMedalsForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDMedalsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDMedalsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
