package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDOpportunitiesInternalServerErrorBody get_characters_character_id_opportunities_internal_server_error
//
// Internal server error
// swagger:model getCharactersCharacterIdOpportunitiesInternalServerErrorBody
type GetCharactersCharacterIDOpportunitiesInternalServerErrorBody struct {

	// get_characters_character_id_opportunities_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id opportunities internal server error body
func (m *GetCharactersCharacterIDOpportunitiesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDOpportunitiesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDOpportunitiesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDOpportunitiesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
