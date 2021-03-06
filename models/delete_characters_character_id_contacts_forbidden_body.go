package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteCharactersCharacterIDContactsForbiddenBody delete_characters_character_id_contacts_forbidden
//
// Forbidden
// swagger:model deleteCharactersCharacterIdContactsForbiddenBody
type DeleteCharactersCharacterIDContactsForbiddenBody struct {

	// delete_characters_character_id_contacts_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete characters character Id contacts forbidden body
func (m *DeleteCharactersCharacterIDContactsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteCharactersCharacterIDContactsForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteCharactersCharacterIDContactsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DeleteCharactersCharacterIDContactsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
