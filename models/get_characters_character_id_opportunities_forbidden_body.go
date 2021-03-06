package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDOpportunitiesForbiddenBody get_characters_character_id_opportunities_forbidden
//
// Forbidden
// swagger:model getCharactersCharacterIdOpportunitiesForbiddenBody
type GetCharactersCharacterIDOpportunitiesForbiddenBody struct {

	// get_characters_character_id_opportunities_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id opportunities forbidden body
func (m *GetCharactersCharacterIDOpportunitiesForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDOpportunitiesForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDOpportunitiesForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDOpportunitiesForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
