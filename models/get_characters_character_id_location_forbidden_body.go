package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDLocationForbiddenBody get_characters_character_id_location_forbidden
//
// Forbidden
// swagger:model getCharactersCharacterIdLocationForbiddenBody
type GetCharactersCharacterIDLocationForbiddenBody struct {

	// get_characters_character_id_location_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id location forbidden body
func (m *GetCharactersCharacterIDLocationForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDLocationForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDLocationForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDLocationForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
