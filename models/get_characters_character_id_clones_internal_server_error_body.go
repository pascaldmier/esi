package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDClonesInternalServerErrorBody get_characters_character_id_clones_internal_server_error
//
// Internal server error
// swagger:model getCharactersCharacterIdClonesInternalServerErrorBody
type GetCharactersCharacterIDClonesInternalServerErrorBody struct {

	// get_characters_character_id_clones_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id clones internal server error body
func (m *GetCharactersCharacterIDClonesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDClonesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDClonesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDClonesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
