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

// GetCharactersCharacterIDMailListsOKBody get_characters_character_id_mail_lists_ok
//
// 200 ok array
// swagger:model getCharactersCharacterIdMailListsOKBody
type GetCharactersCharacterIDMailListsOKBody []*GetCharactersCharacterIDMailListsOKBodyItems

// Validate validates this get characters character Id mail lists o k body
func (m GetCharactersCharacterIDMailListsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	iGetCharactersCharacterIDMailListsOKBodySize := int64(len(m))

	if err := validate.MaxItems("", "body", iGetCharactersCharacterIDMailListsOKBodySize, 1000); err != nil {
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
