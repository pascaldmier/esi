package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDMedalsOKBody get_characters_character_id_medals_ok
//
// 200 ok array
// swagger:model getCharactersCharacterIdMedalsOKBody
type GetCharactersCharacterIDMedalsOKBody []*GetCharactersCharacterIDMedalsOKBodyItems

// Validate validates this get characters character Id medals o k body
func (m GetCharactersCharacterIDMedalsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	iGetCharactersCharacterIDMedalsOKBodySize := int64(len(m))

	if err := validate.MaxItems("", "body", iGetCharactersCharacterIDMedalsOKBodySize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
