package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDSkillsForbiddenBody get_characters_character_id_skills_forbidden
//
// Forbidden
// swagger:model getCharactersCharacterIdSkillsForbiddenBody
type GetCharactersCharacterIDSkillsForbiddenBody struct {

	// get_characters_character_id_skills_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id skills forbidden body
func (m *GetCharactersCharacterIDSkillsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDSkillsForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDSkillsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDSkillsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
