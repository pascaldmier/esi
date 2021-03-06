package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDMailMailIDInternalServerErrorBody get_characters_character_id_mail_mail_id_internal_server_error
//
// Internal server error
// swagger:model getCharactersCharacterIdMailMailIdInternalServerErrorBody
type GetCharactersCharacterIDMailMailIDInternalServerErrorBody struct {

	// get_characters_character_id_mail_mail_id_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id mail mail Id internal server error body
func (m *GetCharactersCharacterIDMailMailIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDMailMailIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDMailMailIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDMailMailIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
