package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCharactersCharacterIDCspaParamsBody post_characters_character_id_cspa_characters
//
// characters object
// swagger:model postCharactersCharacterIdCspaParamsBody
type PostCharactersCharacterIDCspaParamsBody struct {

	// post_characters_character_id_cspa_characters
	//
	// characters array
	// Required: true
	// Max Items: 100
	// Min Items: 1
	// Unique: true
	Characters []int32 `json:"characters"`
}

// Validate validates this post characters character Id cspa params body
func (m *PostCharactersCharacterIDCspaParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharacters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCharactersCharacterIDCspaParamsBody) validateCharacters(formats strfmt.Registry) error {

	if err := validate.Required("characters", "body", m.Characters); err != nil {
		return err
	}

	iCharactersSize := int64(len(m.Characters))

	if err := validate.MinItems("characters", "body", iCharactersSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("characters", "body", iCharactersSize, 100); err != nil {
		return err
	}

	if err := validate.UniqueItems("characters", "body", m.Characters); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersCharacterIDCspaParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersCharacterIDCspaParamsBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDCspaParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
