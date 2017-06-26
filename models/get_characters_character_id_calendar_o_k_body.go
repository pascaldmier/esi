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

// GetCharactersCharacterIDCalendarOKBody get_characters_character_id_calendar_ok
//
// Up to 50 events from now or the event you requested
// swagger:model getCharactersCharacterIdCalendarOKBody
type GetCharactersCharacterIDCalendarOKBody []*GetCharactersCharacterIDCalendarOKBodyItems

// Validate validates this get characters character Id calendar o k body
func (m GetCharactersCharacterIDCalendarOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	iGetCharactersCharacterIDCalendarOKBodySize := int64(len(m))

	if err := validate.MaxItems("", "body", iGetCharactersCharacterIDCalendarOKBodySize, 50); err != nil {
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
