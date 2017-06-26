package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostCharactersCharacterIDCspaForbiddenBody post_characters_character_id_cspa_forbidden
//
// Forbidden
// swagger:model postCharactersCharacterIdCspaForbiddenBody
type PostCharactersCharacterIDCspaForbiddenBody struct {

	// post_characters_character_id_cspa_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this post characters character Id cspa forbidden body
func (m *PostCharactersCharacterIDCspaForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersCharacterIDCspaForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersCharacterIDCspaForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDCspaForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
