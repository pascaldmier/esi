package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCharactersCharacterIDFittingsCreatedBody post_characters_character_id_fittings_created
//
// 201 created object
// swagger:model postCharactersCharacterIdFittingsCreatedBody
type PostCharactersCharacterIDFittingsCreatedBody struct {

	// post_characters_character_id_fittings_fitting_id
	//
	// fitting_id integer
	// Required: true
	FittingID *int32 `json:"fitting_id"`
}

// Validate validates this post characters character Id fittings created body
func (m *PostCharactersCharacterIDFittingsCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFittingID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCharactersCharacterIDFittingsCreatedBody) validateFittingID(formats strfmt.Registry) error {

	if err := validate.Required("fitting_id", "body", m.FittingID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersCharacterIDFittingsCreatedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersCharacterIDFittingsCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDFittingsCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
