package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostCharactersCharacterIDFittingsForbiddenBody post_characters_character_id_fittings_forbidden
//
// Forbidden
// swagger:model postCharactersCharacterIdFittingsForbiddenBody
type PostCharactersCharacterIDFittingsForbiddenBody struct {

	// post_characters_character_id_fittings_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this post characters character Id fittings forbidden body
func (m *PostCharactersCharacterIDFittingsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersCharacterIDFittingsForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersCharacterIDFittingsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDFittingsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
